package main

import "fmt"

func CalmDown() {
	recover()
}

func FreakOut() {
	defer CalmDown()

	panic("Freaking out")
	fmt.Println("It will not be run")
}

func main() {
	FreakOut()
	fmt.Println("Exiting normally")
}
