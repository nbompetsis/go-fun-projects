package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	// runes
	fmt.Println('A') // 65

	num := 1.0000001
	fmt.Printf("Type of num %T\n", num) //float

	fmt.Println(reflect.TypeOf(25))      // int
	fmt.Println(reflect.TypeOf(true))    // bool
	fmt.Println(reflect.TypeOf(5.2))     // float64
	fmt.Println(reflect.TypeOf(1))       // int
	fmt.Println(reflect.TypeOf(false))   // bool
	fmt.Println(reflect.TypeOf(1.0))     // float64
	fmt.Println(reflect.TypeOf("hello")) // string

	var str string = "test"

	stringSize := len(str) + int(unsafe.Sizeof(str))
	fmt.Println("stringSize", stringSize)

	a := [...]string{"t", "e", "s", "t"}
	var b [len(a)]byte
	c := strings.Join(a[:], "")
	copy(b[:], c)

	fmt.Printf("%#v\n", b)
}
