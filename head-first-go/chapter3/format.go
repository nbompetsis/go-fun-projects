package main

import "fmt"

func main() {

	fmt.Printf("A float: %f\n", 3.1475)
	fmt.Printf("An integer %d\n", 10)
	fmt.Printf("A string %s\n", "text")
	fmt.Printf("A boolean %t\n", 3 > 1)

	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Values: %T %T %T\n", 1.2, "\t", true)

	fmt.Printf("Percentage symbol %%\n")
	fmt.Printf("%f liters needed\n", 1.8199999999999998)
	fmt.Println(1.8199999999999998, "liters needed ")

	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5) // padding in decimal value
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Printf("A float: %5.3f\n", 1.111)
}
