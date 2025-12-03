# Go (Golang) Learning Repository

This repository contains my project files and exercises from following Stephen Grider's course on Udemy: **Go: The Complete Developer's Guide (Golang)**.

It serves as a personal log of my progress and a collection of practical examples covering various Go concepts.

## Folder Contents

### `/channels`

This directory contains a detailed exploration of Go's concurrency model. The `main.go` file demonstrates:

*   **Goroutines**: How to create lightweight threads of execution for concurrent tasks.
*   **Channels**: How to use channels for communication and synchronization between goroutines.
*   **Continuous Monitoring**: An implementation of a website status checker that continuously re-checks links in an endless loop.
*   **Function Literals**: Use of anonymous functions (closures) to manage concurrent tasks with delays.

The code is heavily commented to serve as a learning resource for understanding these core Go features.

### `/exercises/OpenFile`

This project is a simple command-line tool that demonstrates file I/O in Go. The `main.go` file shows how to:

*   Read command-line arguments using the `os` package to get a filename.
*   Open and read a file specified by the user.
*   Handle potential errors during file operations.
*   Use `io.Copy` to efficiently stream file contents to standard output (`os.Stdout`), which is more memory-efficient for large files than reading the entire file into memory.

---

*This repository is for educational purposes.*