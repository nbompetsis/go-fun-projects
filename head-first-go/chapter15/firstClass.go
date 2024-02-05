package main

import "fmt"

func sayHello() {
	fmt.Println("sayHello")
}

func sayBye() {
	fmt.Println("sayBye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	twice(sayHello)
	twice(sayBye)
}
