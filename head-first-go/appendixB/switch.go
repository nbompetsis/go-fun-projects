package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomFunc() {
	result := rand.Intn(3) + 1
	switch result {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		panic("invalid number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randomFunc()
}
