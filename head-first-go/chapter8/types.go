package main

import "fmt"

type Part struct {
	description string
	count       int
}

type Car struct {
	name     string
	topSpeed float64
}

func minimumOrder(description string) Part {
	var newPart Part
	newPart.description = description
	newPart.count = 1

	return newPart
}

func showCarInfo(car Car) {
	fmt.Println("Car name", car.name)
	fmt.Println("TopSpeed ", car.topSpeed)
}

func updateCarTopSpeed(car *Car) {
	car.topSpeed = 0
}

func main() {
	var bolts Part = minimumOrder("Hex bolts")
	fmt.Printf("%#v\n", bolts)

	var porsche Car
	porsche.name = "Porsche"
	porsche.topSpeed = 300
	showCarInfo(porsche)

	updateCarTopSpeed(&porsche)
	showCarInfo(porsche)
}
