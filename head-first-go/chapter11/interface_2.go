package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffePot string

func (f CoffePot) TurnOn() {
	fmt.Println("Powering up")
}

func (f CoffePot) Brew() {
	fmt.Println("Heating up")
}

func main() {
	var device Appliance
	device = Fan("MyFan")
	device.TurnOn()

	device = CoffePot("MyPot")
	device.TurnOn()
	// device.Brew()
}
