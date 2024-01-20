package main

import "fmt"

var a = "a0"

func main() {
	a = "a1"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		// fmt.Println(d) // compilation error
	}
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c) // compilation error
	// fmt.Println(d) // compilation error
}
