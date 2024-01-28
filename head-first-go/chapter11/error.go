package main

import (
	"fmt"
	"log"
)

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!\n", o)
}

func CheckTemperature(actual float64, safe float64) error {
	diff := actual - safe
	if diff > 0 {
		return OverheatError(diff)
	}

	fmt.Println("The temperature is ok")
	return nil
}

func main() {
	err := CheckTemperature(121.379, 100.0)

	if err != nil {
		log.Fatal(err)
	}
}
