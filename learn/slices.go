package main

import "fmt"


func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)


	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("add:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// slice literals
	q := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(q)

	r := []bool{true, false, true, false, true, true}
	fmt.Println(r)

	s1 := []struct {
			i int
			b bool
	}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
	}

	fmt.Println(s1)

	// slice default
	s2 := []int{2, 3, 4, 5, 6,7}
	fmt.Println(s2[1:4])
	fmt.Println(s2[:2])
	fmt.Println(s2[1:])

	// slice len and capacity
	printSlice(s2[:0])
	printSlice(s2[:4])
	printSlice(s2[2:])

	// creating slice with make

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
