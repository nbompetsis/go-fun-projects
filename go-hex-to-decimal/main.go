package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

var hex = [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
var hexToDec = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15}
var wg sync.WaitGroup

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please type an decimal input to convert it as hex and vice versa. Run program like the following example\n\" go run main.go 7562 \"")
		os.Exit(1)
	}
	inputAsDecimal, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Invalid decimal input")
		os.Exit(1)
	}

	ch := make(chan string)
	wg.Add(2)

	go ToHex(ch, &wg, inputAsDecimal)
	go ToDecimal(ch, &wg)

	wg.Wait()
}

func ToHex(ch chan string, wg *sync.WaitGroup, number int64) {
	defer wg.Done()

	var result string
	inputDecimal := number
	for inputDecimal != 0 {
		devider, remaining := ToHexDigit(inputDecimal)
		result += hex[remaining]
		inputDecimal = devider
	}
	res := Reverse(result)
	fmt.Println("Hex result of number:", number, "is ->", res)
	ch <- res
}

func ToHexDigit(number int64) (devider int64, remaining int64) {
	devider = int64(number) / int64(len(hex))
	remaining = int64(number) % int64(len(hex))

	return devider, remaining
}

func ToDecimal(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	var result float64
	hexFromCh := <-ch
	revertHex := []rune(Reverse(hexFromCh))
	for exp := len(revertHex) - 1; exp >= 0; exp-- {
		result += MultiplyPowerByNumber(hexToDec[string(revertHex[exp])], exp)
	}
	fmt.Printf("Decimal result of hex: %s is -> %.0f\n", hexFromCh, result)
}

func MultiplyPowerByNumber(num int, exponential int) float64 {
	base := float64(len(hex))
	res := math.Pow(base, float64(exponential))

	return res * float64(num)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
