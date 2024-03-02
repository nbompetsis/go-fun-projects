package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Page  string   `json:"page"`
	Words []string `json:"words"`
}

var TestBook = []Book{
	{
		Page:  "1",
		Words: []string{"a", "b"},
	},
	{
		Page:  "2",
		Words: []string{"c", "d"},
	},
}

func getBooks(c *gin.Context) {
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
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}
