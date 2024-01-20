package main

import (
	"fmt"
	"reflect"
)

func main() {

	var myInt int
	fmt.Println("type of variable", reflect.TypeOf(&myInt))

	var myFloat float64
	fmt.Println("type of variable", reflect.TypeOf(&myFloat))

	var myBool bool
	fmt.Println("type of variable", reflect.TypeOf(&myBool))

	myInt = 10
	var myIntPointer *int = &myInt
	fmt.Println("Pointer value", *myIntPointer)

	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println("Pointer value", *myFloatPointer)

	myInt = 4
	fmt.Println(myInt)
	fmt.Println(&myInt)
	myIntPointer = &myInt
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myIntPointer)
	fmt.Println(myInt)
}
