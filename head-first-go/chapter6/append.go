package main

import "fmt"

func main() {

	var slice = make([]string, 1)
	slice[0] = "first"
	fmt.Println(slice)

	slice = append(slice, "second")
	fmt.Println(slice)

	slice = append(slice, "three", "four", "five")
	fmt.Println(slice, len(slice))

}
