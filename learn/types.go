package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

// numeric constant
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x*0.1
}


func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)

	// types conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	// type inference
	// when declaring a variable without an explicit type (by := operator), the variable's
	// type is inferred from the value on the right hand side
	v := 42  //change me!
	fmt.Printf("v is of type: %T\n", v)

	// constants
	// constants cannot be declared using the := operator
	// constants can be character, string, boolean or numeric
	const Pi = 3.14
	fmt.Println(Pi)

	// numeric constant
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
