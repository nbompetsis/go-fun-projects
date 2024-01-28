package main

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("MethodWithParameters called with:", f)
}

func (m MyType) MethodWithReturnValue() string {
	fmt.Println("MethodWithReturnValue called")
	return "Hi from MethodWithReturnValue method"
}

func main() {
	var value MyInterface
	value = MyType(1)

	value.MethodWithoutParameters()
	value.MethodWithParameters(113.21)
	value.MethodWithReturnValue()
}
