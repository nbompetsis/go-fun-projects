package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) int {
	defer recover()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyAsString := string(body)
	fmt.Println("Size:", len(bodyAsString))
	return len(bodyAsString)
}

func main() {
	go responseSize("https://example.com/")
	go responseSize("https://golang.org")
	time.Sleep(5 * time.Second)
}
