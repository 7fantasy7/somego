package main

import (
	"fmt"
	"sync"
	"time"
)

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

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
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// // Mutex
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	/////
	go say("world")
	say("hello")

	/////
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	/////
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	/////
	fc := make(chan int, 10)
	go fibonacci(cap(fc), fc)
	for i := range fc {
		fmt.Println(i)
	}

	////
	fcs := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-fcs)
		}
		quit <- 0
	}()
	fibonacciSelect(fcs, quit)

	//// Default selection
	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick.")
	//	case <-boom:
	//		fmt.Println("BOOM!")
	//		return
	//	default:
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	/// Mutex
	sc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(sc.Value("somekey"))
}
