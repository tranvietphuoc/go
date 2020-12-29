package main

import (
		"fmt"
		"math"
)

func pow(x, n, lim float64) float64 {
		if v:= math.Pow(x, n); v < lim {
				return v
		}
		return lim
}

func main() {

	// for - while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for forever - break
	for {
		fmt.Println("loop")
		break
	}

	// for continue
	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
	)
}
