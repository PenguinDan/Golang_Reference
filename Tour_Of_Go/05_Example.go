package main

import (
	"fmt"
	"sync"
	"time"
)

// Functions
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // Send sum to c
}

func fibonacii(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// Sends data to the channel
		c <- x
		x, y = y, x+y
	}
	// Closes the integer channel
	close(c)
}

func selectFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// The case that sends a value to the c channel
		case c <- x:
			x, y = y, x+y
		// If quit channel contains a value, this will run
		case <-quit:
			fmt.Println("quit")
			return
		// Default case that runs if no other case is ready
		default:
			fmt.Println("Default case")
			time.Sleep(50 * time.Millisecond)
		}

	}
}

// Using Mutex for Synchronization locks so that a variable can be
// accessed one at a time
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime
	// The following starts a goroutine running the specified function
	go say("world")

	// Using channels
	// Channels are typed conduits through which you can send and receive values with the channel operator <-
	s := []int{7, 2, 8, 1, 2, 1}
	c := make(chan int)     // Creates a channel
	go sum(s[:len(s)/2], c) // Starts a goroutine
	go sum(s[len(s)/2:], c) // Starts another goroutine
	x, y := <-c, <-c        // Gives whatever value in the channel, kind of like an array, passes the value in order
	fmt.Println(x, y, x+y)

	// Buffered channels
	// Sends to a buffered channel block only when the buffer is full
	// Receives block when the buffer is empty
	ch := make(chan int, 2) // Channel that accepts integers with a buffer size of 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Channels with Range and Close
	ct := make(chan int, 10)
	go fibonacii(cap(ct), ct)
	// Prints all of the values in the channel until it is closed
	// Only the sender should close the channel, never the receiver since
	// sending on a closed channel will cause a panic
	// Channels don't need to be closed, closing is only necessary when the receiver
	// must be told there are no more values coming, such as to terminate a range loop
	for i := range c {
		fmt.Println(i)
	}

	// Select goroutines
	// A select statement lets a goroutine wait on multiple communication operations
	// A select blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if multiple are ready
	ct = make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ct)
		}
		quit <- 0
	}()
	// The goroutine above is already blocking until a value in ct is given
	// the channel ct starts being filled in once the bottom method runs
	selectFibonacci(ct, quit)

	// Synchronized Go Routines
	ci := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go ci.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(ci.Value("somekey"))
}
