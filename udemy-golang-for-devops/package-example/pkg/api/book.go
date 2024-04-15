package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func doBookRequest(client http.Client, req string) ([]Book, error) {
	response, err := client.Get(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close() //import step! You need to close the reader

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var book []Book
	if err := json.Unmarshal(body, &book); err != nil {
		fmt.Printf("Something wrong: %v\n", err)
		return nil, err
	}

	return book, nil
}
