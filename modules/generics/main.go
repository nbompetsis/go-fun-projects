package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(data map[string]int64) int64 {
	var sum int64
	for _, value := range data {
		sum += value
	}

	return sum
}

func SumFloats(data map[string]float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}

	return sum
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}

	return sum
}

func SumNumber[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}

	return sum
}

func main() {

	dataInt := map[string]int64{
		"first":  1,
		"second": 100,
	}

	dataFloat := map[string]float64{
		"first":  1.111111111,
		"second": 101.11,
	}

	fmt.Printf("Generic SumInts: %v and SumFloats %v\n",
		SumInts(dataInt),
		SumFloats(dataFloat))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](dataInt),
		SumIntsOrFloats[string, float64](dataFloat))

	fmt.Printf("Generic infered Sums: %v and %v\n",
		SumIntsOrFloats(dataInt),
		SumIntsOrFloats(dataFloat))

	fmt.Printf("Generic Sums Numbers: %v and %v\n",
		SumNumber(dataInt),
		SumNumber(dataFloat))
}
