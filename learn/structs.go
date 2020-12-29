package main

import "fmt"

type Vertex struct {
	X, Y int
}

// struct literal
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y: 0 implicit
	v3 = Vertex{}      // X: 0, Y: 0
	p  = &Vertex{1, 2} //has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	// struct fields
	// v := Vertex{1, 2}
	// v.X = 4  // accessing to X using a dot
	//fmt.Println(v.X)

	// pointer to struct
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)

	// struct literal
	fmt.Println(v1, p, v2, v3)

}
