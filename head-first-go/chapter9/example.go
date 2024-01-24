package main

import "fmt"

type Gallon float64
type Liter float64
type Title string

func main() {
	gallon := Gallon(100)
	fmt.Println(gallon)
	liter := Liter(240)
	fmt.Println(liter)

	gallon = Gallon(Liter(110))
	fmt.Println(gallon)

	carFuel := Gallon(Liter(40.0) * 0.264)
	busFuel := Liter(Gallon(63.0) * 3.785)

	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
	fmt.Println(Liter(1.2) + Liter(3.4))
	fmt.Println(Liter(1.2) * Liter(3.4))
	// fmt.Println(Liter(1.2) + Gallon(3.4)) //not compatible

	fmt.Println(Title("Alien") == Title("Alien"))
	// fmt.Println(Title("Alien") - "2") // not compatible

	fmt.Println(Liter(1.2) + 3.4)
}
