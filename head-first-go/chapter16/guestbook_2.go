package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type GuestBook struct {
	SignaturesCount int
	Signatures      []string
}

func getSignature(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func guestbookHandler(response http.ResponseWriter, request *http.Request) {
	// responseBody := []byte("signature list goes here")
	signatures := getSignature("signatures.txt")
	guestBook := GuestBook{SignaturesCount: len(signatures), Signatures: signatures}

	template, err := template.ParseFiles("view.html")
	check(err)

	// _, err := response.Write(responseBody)
	err = template.Execute(response, guestBook)
	check(err)
}

func newSignature(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("new.html")
	check(err)
	err = template.Execute(response, nil)
	check(err)
}

func createSignature(response http.ResponseWriter, request *http.Request) {
	input := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, input)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(response, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", guestbookHandler)
	http.HandleFunc("/guestbook/new", newSignature)
	http.HandleFunc("/guestbook/create", createSignature)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
