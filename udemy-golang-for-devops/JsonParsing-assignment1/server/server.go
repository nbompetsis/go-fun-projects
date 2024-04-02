package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var words = []string{"one", "two", "three"}
var special = []any{"one", "two", nil}
var extraSpecial = []any{1, 2, "3"}

//	var percentages mapp = Percentages{
//		one:   0.33,
//		three: 0,
//		two:   0.66,
//	}
var percentages map[string]float32 = map[string]float32{
	"one":   0.33,
	"three": 0,
	"two":   0.66,
}

func getParsing(c *gin.Context) {
	fmt.Printf("Percentages: %#v\n", percentages)
	c.JSON(http.StatusOK, gin.H{
		"page":         "assignment1",
		"words":        words,
		"percentages":  percentages,
		"special":      special,
		"extraSpecial": extraSpecial,
	})
}

func main() {
	router := gin.Default()
	router.GET("/parsing", getParsing)
	router.Run("localhost:8080")
}
