package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("unbind error")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	// password check
	if req.Username == "nikos" && req.Password == "bob" {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoibmlrb3MiLCJzdXJuYW1lIjoiYm9iIiwiaWF0IjoxNTE2MjM5MDIyfQ.p6wiTvCJ0NeHkSDLjxYHDLxms8v9flElAMpobaVDYCM"
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error":  "Invalid credentials",
		"status": http.StatusBadRequest,
	})
}
