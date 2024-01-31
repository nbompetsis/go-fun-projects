package main

import "fmt"

func Test1() {
	fmt.Println("Test1")
}

func Test2() {
	defer fmt.Println("End of Test2!")
	defer fmt.Println("End2 of Test2!")
	fmt.Println("Test2")
}

func main() {
	defer fmt.Println("End of main")
	Test1()
	Test2()
}
