package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var quantity int
	// var length, width float64
	// var customerName string

	// quantity = 3
	// length, width = 2, 3
	// customerName = "Nick Bro"

	// fmt.Println(customerName)
	// fmt.Println("has ordered", quantity, "sheets")
	// fmt.Println("each with an area of")
	// fmt.Println(length*width, "square meters")

	// var quantity int = 3
	// var length, width float64 = 1.2, 2.4
	// var customerName string = "Nik Bro"

	// fmt.Println(customerName)
	// fmt.Println("has ordered", quantity, "sheets")
	// fmt.Println("each with an area of")
	// fmt.Println(length*width, "square meters")

	// var quantity = 3
	// var length, width = 1.2, 2.4
	// var customerName = "Nik Bro"
	// fmt.Println(reflect.TypeOf(quantity))
	// fmt.Println(reflect.TypeOf(length))
	// fmt.Println(reflect.TypeOf(width))
	// fmt.Println(reflect.TypeOf(customerName))

	// fmt.Println(customerName)
	// fmt.Println("has ordered", quantity, "sheets")
	// fmt.Println("each with an area of")
	// fmt.Println(length*width, "square meters")

	// Short Variable Declaration.
	quantity := 3
	quantity = 5 // you can assign new value to the variable quantity

	length, width := 1.2, 2.4
	customerName := "Nik Bro"
	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(customerName))

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	// var myFloat float64 = 1.1
	// var myInt int = 5

	// myInt = myFloat // Error
	// fmt.Println("Multiplication of a float and an int is invalid ", myFloat*&myInt) // Error
	// fmt.Println("Is myFloat greater than myInt ? ", myFloat < myInt) // Error

	var myFloat float64
	var myInt int = 5

	myFloat = float64(myInt)
	fmt.Println("MyFloat", myFloat, "and its type is", reflect.TypeOf(myFloat))

	var newLength float64 = 3.75
	var newWidth int = 5
	newWidth = int(newLength)
	fmt.Println("Width", newWidth) // From float to int the fractional part is dropped
}
