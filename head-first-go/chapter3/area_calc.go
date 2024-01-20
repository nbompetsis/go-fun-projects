package main

import (
	"fmt"
	"log"
)

func areaCalc(width float64, height float64) (area float64, err error) {
	if width < 0.0 || height < 0.0 {
		err := fmt.Errorf("One of these params is negative, width: %0.2f or height: %0.2f", width, height)
		// err := errors.New("random error")
		// err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
		// fmt.Println(err.Error())
		// fmt.Println(err)
		return 0, err
	}

	area = width * height
	return (area / 10.0), nil
}

func main() {
	area, err := areaCalc(3.1, 3.8)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("liters needed %0.2f\n", area)

	area, err = areaCalc(4.1, -3.0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("liters needed %0.2f\n", area)
}
