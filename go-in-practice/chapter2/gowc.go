package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	WordsFlag bool `short:"w" long:"words" description:"Print the newline count."`
	LinesFlag bool `short:"l" long:"lines" description:"Print the word counts."`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	_, err := parser.Parse()
	if err != nil {
		fmt.Println("Error parsing command-line arguments:", err)
		os.Exit(1)
	}

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

	if opts.WordsFlag {
		fmt.Println("Words count: ", words)
	}

	if opts.LinesFlag {
		fmt.Println("Lines: ", lines)
	}

	if !opts.WordsFlag && !opts.LinesFlag {
		fmt.Print("use \"./gowc --help\" for more information.")
	}
}
