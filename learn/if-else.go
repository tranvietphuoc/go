package main

import (
	"fmt"
	"math"
)

// if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if with short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// loop and function
func Sqrt(x float64) float64 {
	// calculate the square root of x using Newton-Rapshon method
	z := float64(1)
	for t := 0; t < 10; t++ {
		z -= (z*z - x)/(2*z)
	}
	return z
}


func main() {
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// if
	fmt.Println(sqrt(2), sqrt(-4))

	//if with short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	// calculate squareroot
	fmt.Println(Sqrt(2))
}
