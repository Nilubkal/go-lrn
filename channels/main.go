package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	// Serial implementation
	// for _, link := range links {
	// 	checkLink(link)

	// Serial, async implementation using channels/ go routines
	/*
	  Go Routines:
	  - A goroutine is a lightweight thread of execution. It's not a system thread, but a logical thread that is managed by the Go runtime.
	  - You can think of it as a function that is capable of running concurrently with other functions.
	  - To start a new goroutine, you use the 'go' keyword before a function call. For example: `go myFunction()`.
	  - Goroutines are very cheap to create. They start with a small stack, which can grow and shrink as needed. This means you can have many thousands of them running in a single program.
	  - They run in the same address space, so when you have multiple goroutines accessing shared data, you must synchronize access to prevent race conditions. Channels are one of the ways to achieve this synchronization in Go.
	*/

	/*
	  How Go Routines are Implemented:
	  - The Go runtime includes a scheduler that manages all the goroutines in your program.
	  - This scheduler multiplexes the many goroutines onto a smaller number of OS threads. By default, it creates one OS thread for each available CPU core.
	  - This model allows for concurrency (many tasks making progress) on a single core and true parallelism (many tasks running simultaneously) on multi-core machines.
	  - The scheduler is smart about pausing and resuming goroutines. For example, if a goroutine blocks while waiting for an I/O operation (like an HTTP request), the scheduler will run another goroutine on that OS thread so the CPU doesn't sit idle.
	  - To turn a function call into a goroutine, you simply use the `go` keyword before it. The Go runtime then handles the scheduling of this new goroutine.
	*/
	/*
	  Concurrency vs. Parallelism:
	  - Concurrency is about dealing with lots of things at once. It's a way of structuring your program so that it can work on multiple tasks, but it doesn't mean they are all running at the same instant.
	    On a single CPU core, the Go scheduler will switch between goroutines, giving each a little bit of time to run. This is concurrency.
	  - Parallelism is about doing lots of things at the same time. This requires a machine with multiple CPU cores. The Go runtime can assign different goroutines to run on different cores simultaneously. This is parallelism.
	  - In short: Concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations.
	    Go's concurrency model makes it easy to write programs that can run in parallel on multi-core hardware.
	*/

	/*
	  Main and Child Goroutine Interaction:
	  - When you run `go checkLink(link)` inside the `main` function, you are creating a new "child" goroutine. The `main` function itself runs in a special "main" goroutine.
	  - The `go` keyword is non-blocking. This means the `main` goroutine does not wait for the child goroutine to finish its execution. It immediately continues to the next iteration of the loop.
	  - After the loop finishes, the `main` goroutine has no more code to execute, so it exits.
	  - A critical rule in Go is: when the `main` goroutine finishes, the entire program terminates immediately. It does not wait for any other goroutines to complete.
	  - This is why if you run the code as is, you might see no output or only partial output. The child goroutines are likely terminated before they can finish their HTTP requests and print the results.
	  - To fix this, we need a way for the `main` goroutine to wait for the child goroutines. This is typically done using channels to signal completion.
	  - Channels are used to communicated between different go routines. We need to use a channel to make the main go routine aware when each of the child go routine have completed their execution.
	*/

	/*
	  Channel Structure and Rules:
	  - Channels provide a way for goroutines to communicate and synchronize. Think of them as pipes that connect concurrent goroutines.
	  - **Typed**: Every channel has a type, which is the type of data that can be sent over it. For example, `chan string` is a channel that can only be used to send and receive strings. You cannot send an integer on a string channel.
	  - **Creation**: You create a channel using the `make()` function: `c := make(chan string)`.
	  - **Sending and Receiving**: The `<-` operator is used for both sending and receiving data.
	    - `c <- "hello"`: Sends the string "hello" into the channel `c`.
	    - `msg := <-c`: Receives a value from channel `c` and assigns it to the variable `msg`.
	  - **Blocking**: By default, sends and receives on a channel are blocking.
	    - A send operation will block until another goroutine is ready to receive the value.
	    - A receive operation will block until another goroutine sends a value.
	    - This blocking behavior is what allows goroutines to synchronize their execution without explicit locks.
	*/

	// Creating a channel

	c := make(chan string) // make() is a built-in func that will create a value out of a given type.

	// Concurrent implementation with a channel which will require the checkLink func to receive a second argument of type chan
	for _, link := range links {
		go checkLink(link, c)
	}

	// for {  this is how you do an infinite loop in golang

	// to make it more stylish and not use the endless for
	// we are going to re-write it using <- c == range c
	for l := range c {
		// fmt.Println(<-c)
		// for i := 0; i < len(links); i++ {
		// We ant to loop the execution endlessly with each link re-trying over and over again
		// time.Sleep(5 * time.Second) // this is how you pass a 5 second ( using *5, multiplying it) pause
		// to better implement the wait condition we are going to use a Function literal ( lambda function in Py)
		/*
		  Function Literals:
		  - A function literal is a function without a name. They are also known as anonymous functions or lambdas in other languages.
		  - They are useful for short, one-off functions or when you want to define a function inline without formally declaring it.
		  - **Closure**: Function literals are closures. This means they can access and modify variables from the surrounding scope (the scope in which they are defined).
		  - **Syntax & Example**: You define them using the `func` keyword followed by parameters and a body. You can execute them immediately or assign them to a variable.

		    // Assign to a variable:
		    add := func(a, b int) int {
		        return a + b
		    }
		    // add(2, 3) would return 5

		    // Execute immediately (often used with `go`):
		    go func(message string) {
		        fmt.Println(message)
		    }("Hello from a goroutine!")
		*/

		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}() // we need explicitly to pass () at the end of Function Literals to actually call it/ invoke it.

		// go checkLink(l, c)
		// We receive a value from the channel, which makes the main routine wait.
		// However, we are not assigning it to a variable or printing it.
		// This statement receives the value and immediately discards it, achieving the goal of waiting without producing output.
		// <-c
	}
	// fmt.Println(<-c)
}

// function that will take a link and it will make an http request to it and decide if it responds to it
// adding an argument of type chan and the type of the actual channel makes the func suitable for a goroutine
func checkLink(link string, c chan string) {
	// here we don't care about the response at all (just for the error) therefore -> _
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Error"
		// if we want to continue endlessly the execution with each element we can return in the channel the link element
		c <- link
		return
	}
	// receiving a msg from a channel is usually a blocking operation
	fmt.Println(link, "is up!")
	// c <- "OK"
	// if we want to continue endlessly the execution with each element we can return in the channel the link element
	c <- link
}
