package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var gI, gJ int = 1, 2

const Pi = 3.14

func main() {
	//Types()
	//
	//Structures()
	//
	//Loops()
	//
	//SomeFunc()
	//
	//println(SwitchedDay(2))
	//
	//DeferredFunc()
	//
	//Pointers()
	//
	//FuncAsValue()

	Closure()
}

func Pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(p)  // prints pointer (memory address)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func Loops() {
	// loop
	res1 := ""
	for i := 0; i < 20; i++ {
		res1 += strconv.Itoa(i)
	}
	fmt.Println(res1)

	// omitted
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Structures() {
	// // Structures
	// array
	var arr [10]int // initialized with 0 by default

	fmt.Println("arr: ", arr)
	fmt.Println("len(arr): ", len(arr))
	fmt.Printf("The type of arr is: %T\n", arr)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	// slice
	var slice []int = []int{1, 4, 6, 8, 34, 54, 23, 12, 76, 23}
	fmt.Println("slice: ", slice)
	fmt.Println("len(slice): ", len(slice))
	fmt.Printf("The type of slice is: %T\n", slice)

	// slice through make
	b := make([]int, 0, 5)
	fmt.Println("b: ", b)

	// Slice the slice to give it zero length.
	slice = slice[:0]
	printSlice(slice)

	// Extend its length.
	slice = slice[:4]
	printSlice(slice)

	// Drop its first two values.
	slice = slice[2:]
	printSlice(slice)

	// Append
	slice = append(slice, 10, 11, 12)
	printSlice(slice)

	// loop over slice = range
	for i, v := range slice {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// map
	mmap := make(map[string]int)
	mmap["foo"] = 42
	fmt.Printf("The type of map is: %T\n", mmap)
	fmt.Println(mmap["foo"])

	delete(mmap, "foo")
	// not exists retuns default value (0)
	fmt.Println(mmap["foo"])

	// check existence
	val, ok := mmap["foo"]
	if ok {
		fmt.Printf("The value of foo is: %d\n", val)
	} else {
		fmt.Println("foo not found")
	}

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	pavlo := Person{
		FirstName: "Pavlo",
		LastName:  "Holmes",
		Age:       34,
	}

	fmt.Println(pavlo)

	SliceOfSlices()
}

func SliceOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Types() {
	// multivar
	var c, python, java = true, false, "no!"
	fmt.Println(gI, gJ, c, python, java)

	// // types
	// numeric intX
	intoShort := 127
	fmt.Println("intoShort:", intoShort)
	fmt.Printf("The type of intoShort is: %T\n", intoShort)

	var into int8 = 127
	fmt.Println("into:", into)

	var uinto uint8 = 255
	fmt.Println("uinto:", uinto)

	var floato float32 = 3.14
	fmt.Println("floato:", floato)

	// complex numbers
	var comp1 complex64 = complex(6, 2)
	var comp2 complex64 = 13 + 33i
	fmt.Println(comp1)
	fmt.Println(comp2)

	// Display the type
	fmt.Printf("The type of a is %T and "+
		"the type of b is %T\n", comp1, comp2)
	compSum := comp1 + comp2
	fmt.Println("Comp sum ", compSum)

	// booleans
	x := false
	y := true
	fmt.Println("x && y ", x && y)
	fmt.Println("x || y ", x || y)

	// strings
	str := "Some a bit long, a bit not string"

	fmt.Printf("Length of the string is:%d\n", len(str))
	fmt.Printf("String is: %s\n", str)
	fmt.Printf("Type of str is: %T\n", str)
	fmt.Println("Repeated: ", strings.Repeat(str, 4))

	var char = str[0]
	fmt.Println("Char is: ", char)
	fmt.Printf("Type of char is: %T\n", char)
}

func DeferredFunc() {
	defer println("This is latest")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // deferred order: 9 -> 0
	}

	println("This is first")
}

func SomeFunc() {
	println("hello world from function")
}

func SwitchedDay(number int) string {
	day := ""
	switch number {
	case 1:
		day = "Monday"
	case 2:
		day = "Tuesday"
	default:
		day = "Other"
	}

	return day
}

func FuncAsValue() {
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
