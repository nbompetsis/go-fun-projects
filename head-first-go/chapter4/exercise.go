package main

import (
	"fmt"
	"log"

	"github.com/nbompetsis/gopackages/calc"
	"github.com/nbompetsis/gopackages/keyboard"
)

func main() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(7, 3))
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade > 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
