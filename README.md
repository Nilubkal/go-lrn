# Go (Golang) Learning Repository

This repository contains project files and exercises from Stephen Grider's Udemy course: **Go: The Complete Developer's Guide (Golang)**.

It serves as a personal log of my progress and a collection of practical examples covering various Go concepts.

## Projects

This repository is organized into several projects, each focusing on different aspects of Go.

### 1. Website Status Checker (`/channels`)

This project demonstrates Go's powerful concurrency features through a command-line tool that checks the status of websites.

**Key Concepts:**
*   **Goroutines**: Utilizes lightweight threads for concurrent checking of multiple websites.
*   **Channels**: Implements channels for safe communication and synchronization between goroutines.
*   **Continuous Monitoring**: The application runs in an endless loop to repeatedly check website statuses.
*   **Function Literals**: Employs anonymous functions (closures) to handle the checking logic for each site.

### 2. File Reader CLI (`/exercises/OpenFile`)

This project is a simple command-line tool that showcases file I/O operations in Go. It reads a filename from the command line and prints its contents to the console.

**Key Concepts:**
*   **Command-Line Arguments**: Uses the `os` package to access arguments passed to the program.
*   **File I/O**: Demonstrates how to open, read, and handle files.
*   **Error Handling**: Shows proper error handling for file operations.
*   **Efficient Streaming**: Leverages `io.Copy` to stream file contents directly to standard output (`os.Stdout`), which is highly memory-efficient, especially for large files.

---

*This repository is for educational purposes.*