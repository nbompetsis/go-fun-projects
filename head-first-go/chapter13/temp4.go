package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"page"
)

func responseSize(url string, channel chan page.Page) {
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
	channel <- page.Page{URL: url, Size: len(bodyAsString)}
}

func main() {
	channel := make(chan page.Page)

	urls := []string{"https://example.com", "https://golang.org"}

	for _, url := range urls {
		go responseSize(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		page := <-channel
		fmt.Println("Url:", page.URL, ", size", page.Size)
	}
}
