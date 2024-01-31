package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening file")
	return os.Open(filename)
}

func CloseFile(file *os.File) error {
	fmt.Println("Closing file")
	return file.Close()
}

func GetFloats(filename string) ([]float64, error) {
	var numbers []float64 = []float64{}
	file, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var totalSum float64 = 0.0
	for _, n := range numbers {
		totalSum += n
	}

	fmt.Printf("Total sum: %0.2f\n", totalSum)

}
