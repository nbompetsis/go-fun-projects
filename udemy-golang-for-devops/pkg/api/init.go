package api

import (
	"fmt"
	"log"
	"net/http"
)

type APIInstance interface {
	DoBooksRequest(requestURL string) ([]Book, error)
}

type Options struct {
	Username string
	Password string
	LoginURL string
}

type api struct {
	Options Options
	Client  http.Client
}

func New(o Options) APIInstance {
	api := api{
		Options: o,
		Client: http.Client{
			Transport: &JWTTrasnport{
				Transport: http.DefaultTransport,
				Password:  o.Password,
				Username:  o.Username,
				LoginURL:  o.LoginURL,
			},
		},
	}

	return api
}

func (a api) DoBooksRequest(requestURL string) ([]Book, error) {
	books, err := doBookRequest(a.Client, requestURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%v", books)

	return books, nil
}
