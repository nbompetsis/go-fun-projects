package main

import (
	"fmt"
	"log"

	"github.com/nbompetsis/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.HelloFunction("Nikolas")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)

	names := []string{"Nikolas", "Tom"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for key, value := range messages {
		fmt.Println("[", key, "] has item", value)
	}
}
