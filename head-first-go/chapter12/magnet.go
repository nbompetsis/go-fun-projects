package main

import (
	"fmt"
	"log"
)

func Find(word string, slice []string) bool {
	for _, w := range slice {
		if w == word {
			return true
		}
	}
	return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening Refrigerator")
}

func (r Refrigerator) Close() {
	fmt.Println("Closing Refrigerator")
}

func (r Refrigerator) FindFood(food string) error {
	defer r.Close()
	r.Open()
	ok := Find(food, r)
	if ok {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not Found\n", food)
	}
	return nil
}

func main() {
	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
