package spotify

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
	AlbumType   string  `json:"album_type"`
	Images      []Image `json:"images"`
}

type Track struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	DurationMs   int          `json:"duration_ms"`
	Artists      []Artist     `json:"artists"`
	Album        Album        `json:"album"`
	ExternalURLs ExternalURLs `json:"external_urls"`
	ExternalIDs  ExternalIDs  `json:"external_ids"`
}

// Sous-structs pour les objets imbriqués
type ExternalURLs struct {
	Spotify string `json:"spotify"`
}
type ExternalIDs struct {
	ISRC string `json:"isrc"`
}

type AlbumsPage struct {
	Items []Album `json:"items"`
	Next  string  `json:"next"`
	Total int     `json:"total"`
}
type TracksPage struct {
	Items []Track `json:"items"`
	Next  string  `json:"next"`
	Total int     `json:"total"`
}
