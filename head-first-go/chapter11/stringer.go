package main

import "fmt"

type Gallons float64
type Liters float64
type MilliLiters float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

func (m MilliLiters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func main() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(MilliLiters(12.09248342))
}
