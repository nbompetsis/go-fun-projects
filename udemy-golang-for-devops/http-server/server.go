package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

var TestBook = []Book{
	{
		Page:  "1",
		Input: "drama",
		Words: []string{"a", "b"},
	},
	{
		Page:  "2",
		Input: "novel",
		Words: []string{"c", "d"},
	},
}

func getBooks(c *gin.Context) {
	fmt.Println("Authorization Header %s\n", c.Request.Header.Get("Authorization"))
	c.IndentedJSON(http.StatusOK, TestBook)
}

func postBooks(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		return
	}
	TestBook = append(TestBook, book)
	c.IndentedJSON(http.StatusCreated, book)
}

func main() {
	router := gin.Default()
	router.POST("/login", Login)
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}
