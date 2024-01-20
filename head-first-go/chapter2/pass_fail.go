package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade: ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error on reading input")
	}

	grade, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		log.Fatal("Cannot convert input to float ", err.Error())
	}

	var status string
	if grade >= 60 {
		status = "PASS"
	} else {
		status = "FAIL"
	}

	fmt.Println(status)
}
