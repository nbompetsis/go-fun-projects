package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nbompetsis/album-service-gin/data"
)

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.DataAlbums)
}

func addAlbum(c *gin.Context) {

	var album data.Album

	if err := c.BindJSON(&album); err != nil {
		return
	}

	data.AddAlbum(album)
	c.IndentedJSON(http.StatusCreated, album)
}

func getAlbumById(c *gin.Context) {

	id := c.Param("id")
	album, err := data.FindAlbumById(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Album with id:%v not found", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/album", addAlbum)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}
