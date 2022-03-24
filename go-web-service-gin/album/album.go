package album

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
}

func GetAlbums() []Album {
	return albums
}

func PostAlbum(a Album) {
	albums = append(albums, a)
}
