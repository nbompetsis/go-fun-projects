package main

import "fmt"

type EmptyInterface interface {
}

func AcceptAnything(anything ...EmptyInterface) {
	for _, value := range anything {
		fmt.Printf("Type: %T \n", value)
	}
}

func AcceptAnythingUsingEmptyInterface(anything ...interface{}) {
	for _, value := range anything {
		fmt.Printf("Type: %T \n", value)
	}
}

func main() {
	AcceptAnything("a string", 123, 99.9, true)
	AcceptAnythingUsingEmptyInterface("a string", 123, 99.9, true)
}
