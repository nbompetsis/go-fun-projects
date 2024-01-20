package main

import "fmt"

func main() {

	var notes [7]string
	notes[0] = "ZERO Value"

	for i := 0; i < len(notes); i++ {
		if notes[i] != "" {
			fmt.Println("Element:", i, "->", notes[i])
		}
	}

	notes = [7]string{"do", "re", "mi", "fa", "so"}

	for i := 0; i < len(notes); i++ {
		if notes[i] != "" {
			fmt.Println("Element:", i, "->", notes[i])
		}
	}
	// fmt.Println("Element:", tmp, "->", notes[10]) panic runtime error

	for index, value := range notes {
		if value != "" {
			fmt.Println("Element:", index, "->", value)
		}
	}

	for _, value := range notes {
		if value != "" {
			fmt.Println("Element ->", value)
		}
	}

}
