package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args provides access to raw command-line arguments.
	// The first argument, os.Args[0], is the path to the program itself.
	// os.Args[1] will be the first argument passed by the user.
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
	fileName := os.Args[1]
	// Print all the arguments passed to this script
	fmt.Println("Arguments passed to this script:")
	for i, args := range os.Args {
		fmt.Printf("Argument %d: %s\n", i, args)
	}
	// Read the file using a func with ReadFIle
	// fmt.Printf("Attempting to read file: %s\n", fileName)
	// data, err := readFile(fileName)
	// if err == nil {
	// fmt.Printf("File content:\n%s\n", string(data))
	// }
	// Using the os.Open for more detailed control over a file and memory optimized
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error opening file %s: %v \n", fileName, err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}

func readFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %v \n", fileName, err)
		return nil, err
	}
	return data, nil
}
