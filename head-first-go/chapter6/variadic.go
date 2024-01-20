package main

import "fmt"

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {

	slice := []int{1, 2, 3}

	severalInts(slice...)

	mix(1, true, "test")
}
