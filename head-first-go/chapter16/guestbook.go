package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func guestbookHandler(response http.ResponseWriter, request *http.Request) {
	// responseBody := []byte("signature list goes here")
	template, err := template.ParseFiles("view.html")
	check(err)

	// _, err := response.Write(responseBody)
	err = template.Execute(response, nil)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", guestbookHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
