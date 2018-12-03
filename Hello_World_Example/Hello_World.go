// Package main, executable comman should always use the package name "main"
package main

// Import a standard library for print formatting
import (
	"fmt"

	"github.com/PenguinDan/Golang_Reference/Hello_World_Example/Package/utilities"
)

// Define the entry point of the application
func main() {
	// Simply prints a line to the console using the fmt package
	fmt.Println("Hello, world.")
	// Print the Hello, world. backwards
	backwardsString := utilities.Reverse("Hello, world.")
	fmt.Println(backwardsString)
}

// Build and run the application using the go tool and the following command
// go install github.com/user/Hello_World.go or go install Hello_World.go
// Now the above code will be able to be ran from anywhere
// since the go tool finds the source code by looking for the github.com/user/hello inside GOPATH
// You can type Hello_World from anywhere in your desktop now since ~/Desktop/Golang_Environment/bin
// has been added to the PATH environment variable
