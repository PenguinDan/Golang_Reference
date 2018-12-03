package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

// ---- Defining Type methods --------
type Vertex struct {
	X, Y float64
}

// The first parameter value is the type that we are attaching this method to
// and using its values
// Below can be used as v.Abs()
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Below can be used as Abs(v)
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer Type method receivers allow us to change the values of the types
// Pointer receivers are also faster since we dont need to copy the values
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Type methods are not exclusive to Structs
type MyFloat float64

// Below can be used as f.Abs()
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ------------- Defining Interfaces -----------
type Abser interface {
	Abs() float64
}

// A type implements an interface by implementing its methods. There is no
// explicit declaration of intent, no "implements" keyword
type I interface {
	M()
}
type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicity declare that it does so
func (t T) M() {
	fmt.Println(t.S)
}

// Empty interfaces that sepcifies zero methods is known as the "empty interface"
// An empty interface may hold values of any type since every type implements atleast 0 methods
// Empty intefaces are used by code that handles values of unknown type
// For example : interface{}

// Type Switches with intefaces
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// ----------------- Error Handling ---------------
type MyError struct {
	When time.Time
	What string
}

// Define a MyError Type method that returns a string
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// Simple function that returns an error
func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

// -------------- IO ----------------------------
// The io.Reader interface has a Read method:
// Input is the byte array and returns an interger n and error err
//func (T) Read(b []byte) (n int, err error)

func main() {
	// Vertex methods
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// Go automatically does (&v).Scale(10) instead
	v.Scale(10)
	fmt.Println(v)
	// MyFloat methods
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Defining Interfaces
	var a Abser                // Creates an instance of an interface
	f2 := MyFloat(-math.Sqrt2) // Creates the MyFloat Type object
	a = f2                     // We can do this because MyFloat has a method that implements the Abs signature
	fmt.Println(a.Abs())
	v2 := Vertex{3, 4} // Creates the Vertex Type object
	a = &v2            // Vertex's version of the Abs() method uses a pointer
	fmt.Println(a.Abs())
	// Implicit interfaces
	var i I = T{"Hello"}
	i.M()

	// Empty Interfaces
	var i interface{}
	i = 42
	i = "hello"

	// Type assertions with empty interfaces
	var it interface{} = "hello"
	s := i.(string)
	fmt.Println(s) // Prints "Hello"
	s, ok := i.(string)
	fmt.Println(s, ok) // Prints "Hello" True
	ft, ok := i.(float64)
	fmt.Println(f, ok) // Prints 0 False
	ft = i.(float64)   // Panic state
	fmt.Println(ft)

	// Type switches
	do(21)      // Prints Twice 21 is 42
	do("Hello") // Prints "hello" is 5 bytes long
	do(true)    // Prints I don't know about type bool

	// Run the error case
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Implementing the IO
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b) // Strings implements the Read interface, modifies b and outputs the number of lines input and and error if an error occurred
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	// The above prints the following
	/*
		n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
		b[:n] = "Hello, R"
		n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
		b[:n] = "eader!"
		n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
		b[:n] = ""
	*/
}
