package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	PI, THETA = 3.14, 1
)

type Book struct {
	Page  string   `json:"page"`
	Words []string `json:"words"`
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: .http-server <url>")
	}

	urlParam := os.Args[1]
	if _, err := url.ParseRequestURI(urlParam); err != nil {
		log.Fatalf("Url is invalid, url: %v\n", os.Args[1])
	}
	response, err := http.Get(urlParam)
	if err != nil {
		log.Fatalf("Something wrong: %v\n", err)
	}

	defer response.Body.Close() //import step! You need to close the reader

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Something wrong: %v\n", err)
	}

	fmt.Printf("Http status %d, body %s", response.StatusCode, string(body))

	var books []Book

	if err := json.Unmarshal(body, &books); err != nil {
		log.Fatalf("Something wrong: %v\n", err)
	}

	fmt.Printf("This is how to unmarshal a json response: %v \n", books)

}
