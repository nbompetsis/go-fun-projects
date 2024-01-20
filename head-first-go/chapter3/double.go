package main

import (
	"fmt"
)

func double(number *int) {
	*number *= 2
}

func main() {
	var myInt int = 10
	double(&myInt)
	fmt.Println(myInt)
}
