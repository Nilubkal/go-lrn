package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)

	}
	// To get the body of the response we need to dive to a something like a rabbit hole:
	//	Response struct
	//		Status string
	// 		StatusCode int
	//		Body io.ReadCloser  -> ( field can be an interface, therefore you can put any type in here as long as satisfies the interface)
	// io.ReadCloser interface
	//	    Reader interface      -->   Read ([]byte) (int, error) - represents the number of bytes read from the interface more like a statistical value not real data.
	//      Closer interface 	  --> 	Close() (error)
	// In Go you can take multiple interfaces and assemble them together to form a new interface.

	// --- A more detailed explanation ---
	// The 'resp' variable holds the server's response. A key part of it is the 'Body'.
	//
	// The 'Body' isn't the raw data (like a string of HTML). Instead, its type is 'io.ReadCloser'.
	// 'io.ReadCloser' is an interface, which is like a contract in Go. It says: "any data
	// type that wants to be a response body must know how to do two things":
	//
	// 1. Read(): It must have a 'Read' method to let us read the body's data, often in
	//    small chunks. This is useful for large responses, as we don't have to load
	//    everything into memory at once. This behavior comes from the 'io.Reader' interface.
	//
	// 2. Close(): It must have a 'Close' method. After we're done reading the body, we
	//    MUST close it to release the underlying network connection. Forgetting to close
	//    the body is a common source of bugs and resource leaks. This behavior comes
	//    from the 'io.Closer' interface.
	//
	// Go allows combining simple interfaces (like Reader and Closer) to create more
	// descriptive ones (like ReadCloser). This makes the code flexible and efficient.

	// fmt.Println(resp)

	// make is invoked to create a []byte slice with length of 99999 elements (the elements are empty and pre-created by using make).
	// the reason to use a pre-defined length of the slice is because the Read interface is not meant to resize the slice if it becomes full

	// This is why its using such a large size of 99999 elements !
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// whenever there is a byte slice it can be easily turned into a string by calling the string() func
	// fmt.Println(string(bs))

	// Some advanced representation of the same output but in a more direct GO style way
	//
	// The Writer interface is the counterpart to the Reader; it writes data from a byte slice out to a destination.
	// Where Reader reads *into* a slice, Writer writes *from* a slice, and `os.Stdout` is a Writer for the console.
	// The io.Copy function automatically handles reading all data from the Reader and writing it to the Writer.
	// func Copy (dst Writer, src Reader) (written int64, err error) - definition of the function.
	// io.Copy(os.Stdout, resp.Body)

	// We are going to implement our own Writer interface to show how this is done
	// from the docs we now that a Writer type is represented like this
	// type Writer interface {
	//     Write(p []byte) (n int, err error)
	// }

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
