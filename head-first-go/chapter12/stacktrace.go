package main

func four() {
	three()
}

func three() {
	second()
}

func second() {
	one()
}

func one() {
	panic("Throwing from method one")
}

func main() {
	four()
}
