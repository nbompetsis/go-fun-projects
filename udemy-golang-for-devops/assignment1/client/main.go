package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ResponseBody struct {
	Page         string             `json:"page"`
	Words        []string           `json:"words"`
	Percentages  map[string]float32 `json:"percentages"`
	Special      []interface{}      `json:"special"`
	ExtraSpecial []interface{}      `json:"extraSpecial"`
}

func main() {

	resp, err := http.Get("http://localhost:8080/parsing")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response: ")
	fmt.Println("\"page\":", "\""+responseBody.Page+"\"")
	fmt.Println("\"words\":", responseBody.Words)
	fmt.Println("\"percentages\":", responseBody.Percentages)
	fmt.Printf("\"special\": %v\n", responseBody.Special)
	fmt.Printf("\"extraSpecial\": %v\n", responseBody.ExtraSpecial)

}
