// Package main, executable packages always begin with main
package main

import (
	"fmt"
	"math"
)

// ------------------ Example Functions -------------------
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// --------------------------- Structs ------------------------
// A structure is a collection of fields
type Vertex struct {
	X int
	Y int
}
type Vertex2 struct {
	X, Y int
}
type Vertex3 struct {
	Lat, Long float64
}

// Declaring Structs
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2} // Has type *Vertex
)

// ----------------- Function Values and Parameters -------------------
// Takes in a function that takes in 2 float64 values as arguments and returns a float64
// The compute function then returns the float64 value itself
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function Closures
// A closure is a function value that references variables and keeps those variables alive
// for continuous access
// The following is a closure that returns a function, each call of adder() has its own sum
// variable that it keeps and continuously references
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Go has pointers which basically  as the C++ pointers
	// Below is a pointer that holds the memory address of a value
	var integerPointer *int
	i := 42
	integerPointer = &i

	// The * dereferences the pointer
	fmt.Println(*integerPointer)
	*integerPointer = 21

	// ------------- Structs --------------
	fmt.Println(Vertex{1, 2})
	// Declaring a struct
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// ------- Arrays -----------
	// Variable of two strings
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	// Declaring the values of the array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Dynamic arays or Slices
	// a[low : high]
	var s []int = primes[1:4]
	fmt.Println(s)
	// Slices do not store any data, it just describes a section of an underlying array
	// Changing the elements of a slice modifies the corresponding elements of its underlying array
	// Other slices that share the same underlying array will see those changes
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	// Get a slice of the declared names, slice low high values work the same as python
	at := names[0:2]
	bt := names[1:3]
	fmt.Println(at, bt)
	bt[0] = "XXX" // Modifies the original
	fmt.Println(at, bt)
	fmt.Println(names)
	// A slice literal creates an array, then builds a slice that references it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)
	// A slice has both a length and capacity
	// Length : The number of elements it contains
	// Capacity : The number of elements in the underlying array, counting from the first element in the slice
	slice := []int{2, 3, 5, 7, 11, 13}
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	// Extend its length.
	s = s[:4]
	printSlice(s)
	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Make function creates dynamically-sized arrays
	// Allocate a zeroed array and return a slice that refers to the array
	arr := make([]int, 5)
	// To specify a capacity, pass a third argument
	brr := make([]int, 0, 5) // length = 0, capacity = 5
	brr = brr[:cap(brr)]     // length = 5, capacity = 5
	brr = brr[1:]            // length = 4, capacity = 4

	// Appending values to a slice
	var dSlice []int
	printSlice(dSlice)
	// Append works on nil slices
	dSlice = append(dSlice, 0)
	printSlice(dSlice)
	// The slice grows as necessary
	dSlice = append(dSlice, 1)
	printSlice(dSlice)
	// We can add more than one element at a time
	dSlice = append(dSlice, 2, 3, 4, 5)
	printSlice(dSlice)

	// Iterating over a slice, define range <Slice>
	// Iterating over a slice returns the index and a copy of the element at the index
	var powSlice = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range powSlice {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// If you only want the index, drop the ", value" altogether
	for i := range powSlice {
		// 2**i
		powSlice[i] = 1 << uint(i)
	}
	// You can skip either the index or value by assigning to _
	for _, value := range powSlice {
		fmt.Printf("%d\n", value)
	}

	// --------------- Maps ------------------
	var m map[string]Vertex3
	// Get a map using the make function
	m = make(map[string]Vertex3)
	m["Bell Labs"] = Vertex3{
		40.68433, -74.3997,
	}
	fmt.Println(m["Bell Labs"])
	// Defining a map literraly
	var mt = map[string]Vertex3{
		"Bell Labs": Vertex3{
			40.7, -80.1,
		},
		"Google": Vertex3{
			111, 90.2,
		},
	}
	var mp = map[string]Vertex3{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	// Insert or update an element in the map
	mp["New Key"] = Vertex3{-13, 12}
	// Retrieve an element from the map
	elem := mp["New Key"]
	// Delete an element from the map
	delete(mp, "New Key")
	// Test that a key is present, ok = True if present, false otherwise and elem will be 0
	elem, ok := mp["New Key"]

	// Functions as variables
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Function Closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
