package main

import "fmt"

func SayHello() {
	fmt.Println("sayHello")
}

func Divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var copySayHelloFunc func()
	var copyDivideFunc func(int, int) float64

	copySayHelloFunc = SayHello
	copyDivideFunc = Divide

	copySayHelloFunc()
	fmt.Println(copyDivideFunc(1, 2))
}
