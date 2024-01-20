package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/nbompetsis/gopackages/strings"
)

func main() {

	lines, err := strings.ReadLinesFromFile("votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}
	sort.Strings(counts)
	for k, v := range counts {
		fmt.Println("Name: ", k, ", count: ", v)
	}
}
