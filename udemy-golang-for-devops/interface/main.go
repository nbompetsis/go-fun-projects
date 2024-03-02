package main

import (
	"fmt"
	"io"
	"log"
)

type SlowReader struct {
	Content string
}

func (s *SlowReader) Read(p []byte) (n int, err error) {
	var totalBytes int
	runeOfContent := []byte(s.Content)
	for k := range runeOfContent {
		totalBytes += copy(p[k:k+1], s.Content[k:k+1])
	}

	return totalBytes, io.EOF
}

type QuickReader struct {
	Content string
}

func (s *QuickReader) Read(p []byte) (n int, err error) {
	b := copy(p, []byte(s.Content))
	return b, io.EOF
}

func main() {
	var mySlowReader SlowReader = SlowReader{
		Content: "Hello World, Ελληνικα, 漢字",
	}
	out, err := io.ReadAll(&mySlowReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output : %s\n", out)

	var qReader QuickReader = QuickReader{
		Content: "TEST, Ελληνικα, 漢字",
	}
	out, err = io.ReadAll(&qReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output QuickReader : %s\n", out)
}
