package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/nbompetsis/pkg/api"
)

func main() {

	var (
		username   string
		password   string
		requestURL string
		parsedURL  *url.URL
		err        error
		books      []api.Book
	)

	flag.StringVar(&username, "user", "", "use a user to access our api")
	flag.StringVar(&password, "pass", "", "use a password to access our api")
	flag.StringVar(&requestURL, "url", "", "url to access")

	flag.Parse()

	if username == "" || password == "" {
		fmt.Printf("Please provide username and password")
		flag.Usage()
		os.Exit(1)
	}

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Validation error: URL is not valid: %v", err)
		flag.Usage()
		os.Exit(1)
	}

	apiInstance := api.New(api.Options{Username: username, Password: password, LoginURL: parsedURL.Scheme + "://" + parsedURL.Host + "/login"})
	if books, err = apiInstance.DoBooksRequest(parsedURL.Scheme + "://" + parsedURL.Host + "/books"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
}
