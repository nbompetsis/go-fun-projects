package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var wordsFlag bool
var linesFlag bool

func init() {
	flag.BoolVar(&linesFlag, "lines", false, "Print the newline count")
	flag.BoolVar(&linesFlag, "l", false, "Print the newline count")

	flag.BoolVar(&wordsFlag, "words", false, "Print the word counts.")
	flag.BoolVar(&wordsFlag, "w", false, "Print the word counts.")
}

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	words := 0
	lines := 0

	for scanner.Scan() {
		text := scanner.Text()
		words += len(strings.Fields(text))
		lines++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	if wordsFlag {
		fmt.Println("Words count: ", words)
	}

	if linesFlag {
		fmt.Println("Lines: ", lines)
	}

	if !linesFlag && !wordsFlag {
		fmt.Print("use \"./gowc --help\" for more information.")
	}
}
