package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	var sum float64 = 0
	for _, arg := range arguments {
		res, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += res
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Avg: %0.2f\n", sum/sampleCount)
}
