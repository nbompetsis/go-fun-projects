package main

import (
	"fmt"
	"log"
)

func doSomething(input int) error {
	if input%2 == 0 {
		return nil
	}

	return fmt.Errorf("The %d is not an even number", input)
}

func main() {

	// err := doSomething(2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = doSomething(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// or second approach
	if err := doSomething(2); err != nil {
		log.Fatal(err)
	}

	if err := doSomething(1); err != nil {
		log.Fatal(err)
	}
}
