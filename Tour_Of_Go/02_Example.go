// Package main, executable packages always begin with main
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// ----------------------- Conditionals -----------------------
// Conditional statements do not need () but need to contain the
// surrounding {}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If statements can start with a short statement to execute before the condition
// these variables declared by the statement are only in scope until the end of the if and else statements
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// The main entry point of the executable
func main() {
	// Go contains only one looping construct, the for loop which contains three components
	// separated by semicolons
	/*
		the init statement: executed before the first iteration
		the condition statement: evaluated before every iteration
		the post statement: executed at the end of every iteration
	*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// The init and post statements are optional, the for statement without
	// the two components is Go's version of the while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// Infinite looping while loop
	for {

	}

	// ---------------- Conditionals ------------
	// Print the value from the first conditionals function
	fmt.Println(sqrt(2), sqrt(-4))
	// Print the value from the second conditionals function
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// ------------- Switch Statements -----------
	// Switch statements only run a single case
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s", os)
	}
	// Instead of writing long if-else chains, we can have a switch true block
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good Evening")
	}

	// ------------------------- Defering ----------------
	// A defer statement defers the execution of a funciton until the surrounding function returns
	// meaning, it won't be executed until this main finishes
	defer fmt.Println("World")
	fmt.Println("Hello")
	// Defered statements are in the last-in first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")

}
