package main

import "fmt"

type MyString string

func (str MyString) sayHi() {
	fmt.Println("Say hi from:", str)
}

func main() {

	str := MyString("test")
	str.sayHi()

	anotherStr := MyString("second-test")
	anotherStr.sayHi()

}
