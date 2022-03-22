package main

import (
	"fmt"
	"log"
	"net/http"
)

type ServingReq struct {
	pattern       string
	methodHandler func(w http.ResponseWriter, req *http.Request)
}

var methodNotAllowed = func(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func main() {

	loginRequest := ServingReq{
		pattern: "/login",
		methodHandler: func(w http.ResponseWriter, req *http.Request) {
			if req.Method != "POST" {
				methodNotAllowed(w, req)
				return
			}
			fmt.Fprintf(w, "This is the login page\n")
		},
	}

	aboutRequest := ServingReq{
		pattern: "/about",
		methodHandler: func(w http.ResponseWriter, req *http.Request) {
			if req.Method != "GET" {
				methodNotAllowed(w, req)
				return
			}
			fmt.Fprintf(w, "This is the about page\n")
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc(aboutRequest.pattern, aboutRequest.methodHandler)
	mux.HandleFunc(loginRequest.pattern, loginRequest.methodHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
