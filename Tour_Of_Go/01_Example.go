// Package main, executable packages always begin with main
package main

// This program using the below packages
// By convention the package is the same as the last element of the import path
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// ---------------------- Variables ----------------------------------
// Package level boolean variables with default "false" values
var c, python, java bool

// Variables with initial values, if an initializer is present
// the type can be ommitted
var it, jt int = 1, 2
var ip, jp = 1, 2

//Variables can be grouped in blocks
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Declaring constant variables, they cannot be declared using the := syntax
const Pi = 3.14

// ---------------------- Functions ----------------------------------
func add(x int, y int) int {
	return x + y
}

// If multiple parameters share a type, the type can be omitted until the end
func sub(x, y, z int) int {
	return x - y - z
}

// A function can return any number of values
func swap(x, y string) (string, string) {
	return y, x
}

// Naked functions are able to return named variables which are treated as variables
// defined at the top of the function
// These names should be used to document the meaning of the return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// The main entry point of the executable
func main() {
	// Seed the random function since it is deterministic, meaning it will always
	// spit out the same number given the same seed value
	rand.Seed(int64(time.Now().Second()))
	// Print a random number, 10 is the maximum cap from 0 - 10, I think
	fmt.Println("My favorite number is", rand.Intn(10))

	// Prints the square root of a number, type cast
	fmt.Printf("Square root of %g is %g.\n", float64(7), math.Sqrt(7))

	// All package values and functions that are exposed to an outside environment
	// must begin with a capital letter, thats how they are exposed
	// fmt.Println(math.pi) won't work but below, the following will work:
	fmt.Println(math.Pi)

	// ---------------- Functions -------------------
	// Calls relating to the declared functions
	// Add two values given to the declared function
	fmt.Println(add(42, 34))
	// Subtract 3 values from each other
	fmt.Println(sub(100, 30, 14))
	// Swap the two values
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	// Naked function that returns two values
	fmt.Println(split(17))

	// ----------- Variables ------------------
	var i int
	var ct, pythont, javat = true, false, "no!"
	const Truth = true
	// Print out the values of the declared variables
	fmt.Println(i, c, python, java)
	fmt.Println(it, jt, ct, pythont, javat)
	// Print out the values declared in the declaration block
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// Print out the declared constant variables
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)
}
