package main

import "fmt"

func SayHello() {
	fmt.Println("sayHello")
}

func Divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func Multiply(a int, b int) float64 {
	return float64(a * b)
}

func doMath(theFunction func(int, int) float64) {
	fmt.Println(theFunction(1, 2))
}

func main() {
	doMath(Divide)

	doMath(Multiply)
}
