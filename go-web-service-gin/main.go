package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nbompetsis/web-service-gin/album"
)

func main() {
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.GET("/album/:id", getAlbum)
	r.POST("/album", postAlbum)
	r.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	albums := album.GetAlbums()
	if len(albums) > 0 {
		c.IndentedJSON(http.StatusOK, albums)
		return
	}
	c.IndentedJSON(http.StatusBadRequest, "No Albums")
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")
	albums := album.GetAlbums()
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "Album not found")
}

func postAlbum(c *gin.Context) {
	var newAlbum album.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums := album.GetAlbums()
	for _, a := range albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusBadRequest, "Album has been already imported")
			return
		}
	}
	album.PostAlbum(newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}
