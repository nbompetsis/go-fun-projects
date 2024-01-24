package main

import "fmt"

var student struct {
	name   string
	id     int
	grade  float64
	active bool
}

func main() {
	student.name = "Nikos"
	student.id = 1
	student.grade = 19.1
	student.active = true
	fmt.Printf("%#v\n", student)

}
