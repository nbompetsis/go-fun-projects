package data

import (
	"fmt"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var DataAlbums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func AddAlbum(a Album) {
	DataAlbums = append(DataAlbums, a)
}

func FindAlbumById(id string) (Album, error) {
	for _, value := range DataAlbums {
		if value.ID == id {
			return value, nil
		}
	}

	return Album{}, fmt.Errorf("The id: %v not found", id)
}
