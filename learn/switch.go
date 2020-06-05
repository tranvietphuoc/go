package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	//
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
	//
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i'm a bool")
		case int:
			fmt.Println("i'm an int")
		default:
			fmt.Printf("don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// go run os
	fmt.Println("Go runs on:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd, plan9, windows,...
		fmt.Printf("%s.\n", os)
	}
}
