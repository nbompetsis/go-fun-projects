package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers [3]float64
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = value
		i += 1
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println(numbers)

	var totalNumber float64
	for _, value := range numbers {
		totalNumber += value
	}

	fmt.Println("Final result", totalNumber/float64(len(numbers)))
}
