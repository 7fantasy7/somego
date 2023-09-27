package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

// Pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// todo skipped https://go.dev/tour/methods/6
// Methods and pointer indirection

// Rule:
// In general, all methods on a given type should have either value or pointer receivers,
// but not a mixture of both.

func main() {
	v := Vertex{3, 4}

	v.Scale(10)

	// method
	fmt.Println(v.Abs())

	// function
	fmt.Println(Abs(v))

	// method on non struct type
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	var a Abser
	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
	//
	//// empty interface
	//var i interface{}
	//describe(i)

	// type assertion
	//TypeAssertions()

	//TypeSwitch()
	//
	//p1 := Person{"Arthur Dent", 42}
	//p2 := Person{"Zaphod Beeblebrox", 9001}
	//fmt.Println(p1, p2)
	//
	//// Errors
	//if err := runError(); err != nil {
	//	fmt.Println(err)
	//}

	//DoRead()

	doImage()
}

func DoRead() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func TypeSwitch() {
	do(21)
	do("hello")
	do(true)
}

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

func TypeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Person struct {
	Name string
	Age  int
}

// ToString'er :)
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Errors
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func runError() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func doImage() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
