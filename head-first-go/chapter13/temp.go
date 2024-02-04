package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string) {
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
	// fmt.Println(bodyAsString)
	fmt.Println("Body length:", len(bodyAsString))
}

func main() {
	responseSize("https://example.com/")
	responseSize("https://golang.org")
}
