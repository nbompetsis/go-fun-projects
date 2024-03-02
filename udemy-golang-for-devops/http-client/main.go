package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response interface {
	GetResponse() string
}

type Book struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (b Book) GetResponse() string {
	return strings.Join(b.Words, ", ")
}

func doRequest(req string) (Response, error) {
	if _, err := url.ParseRequestURI(req); err != nil {
		log.Fatalf("Url is invalid, url: %v\n", os.Args[1])
	}
	response, err := http.Get(req)
	if err != nil {
		log.Fatalf("Something wrong: %v\n", err)
	}

	defer response.Body.Close() //import step! You need to close the reader

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var book Book

	if err := json.Unmarshal(body, &book); err != nil {
		log.Fatalf("Something wrong: %v\n", err)
	}

	return book, nil
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: .http-server <url>")
	}

	urlParam := os.Args[1]
	response, err := doRequest(urlParam)
	if err != nil {
		log.Fatalf("Something wrong: %v\n", err)
	}

	fmt.Printf("This is how to unmarshal a json response: %v \n", response)
}
