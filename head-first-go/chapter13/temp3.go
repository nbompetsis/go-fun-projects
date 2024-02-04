package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
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
	channel <- len(bodyAsString)
}

func main() {
	channel := make(chan int)
	go responseSize("https://example.com", channel)
	go responseSize("https://golang.org", channel)

	fmt.Println("Size:", <-channel)
	fmt.Println("Size:", <-channel)
}
