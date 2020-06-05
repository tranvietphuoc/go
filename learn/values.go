package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// variables declared without an explicit inital value are given their zero value.
	// 0: for numeric type
	// false: for boolean type
	// "": for string type
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
