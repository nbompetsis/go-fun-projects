package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, req *http.Request) {
	message := []byte("Hello World!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
