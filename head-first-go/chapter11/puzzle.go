package main

import "fmt"

type Trunk string

func (t *Trunk) Accelerate() {
	fmt.Println("Speeding up")
}

func (t *Trunk) Brake() {
	fmt.Println("Stopping")
}

func (t *Trunk) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t *Trunk) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func TryVehicle(v Vehicle) {
	v.Accelerate()
	v.Steer("left")
	v.Steer("right")
	v.Brake()

	truck, ok := v.(*Trunk)
	if ok {
		truck.LoadCargo("test cargo")
	}
}

func main() {
	truck := Trunk("Fnord 123")
	TryVehicle(&truck)
}
