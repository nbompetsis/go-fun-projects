package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Printf("one\n")
	go testFunction(c)
	<-c
	fmt.Printf("three\n")
}

func testFunction(c chan int) {
	fmt.Printf("two\n")
	c <- 1
}
