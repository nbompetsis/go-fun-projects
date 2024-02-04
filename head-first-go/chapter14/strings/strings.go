package main

import (
	"fmt"
	"strings"
)

func JoinWithCommas(words []string) string {
	if words == nil || len(words) == 0 {
		return ""
	}

	if len(words) == 1 {
		return words[0]
	}

	if len(words) == 2 {
		return strings.Join(words[:], " and ")
	}

	result := strings.Join(words[:len(words)-1], ", ")
	result += ", and "
	result += words[len(words)-1]
	return result
}

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
}
