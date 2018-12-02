package main

// Import a standard library for print formatting
import "fmt"

// Define the entry point of the application
func main() {
	fmt.Println("Hello, world.")
}

// Build and run the application using the go tool and the following command
// go install github.com/user/Hello_World.go or go install Hello_World.go
// Now the above executable will be able to be ran from anywhere
// since the go tool finds the source code by looking for the github.com/user/hello inside GOPATH
