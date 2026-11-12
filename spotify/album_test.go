package spotify

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetArtistAlbums(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				offset := r.URL.Query().Get("offset")
				if offset == "0" {
					fmt.Fprint(w, `{"items":[{"id":"a1","name":"Album 1"}],"next":"suite","total":2}`)
				} else {
					fmt.Fprint(w, `{"items":[{"id":"a2","name":"Album 2"}],"next":null,"total":2}`)
				}
			}))
	defer srv.Close()

	c := NewClient(2 * time.Second)
	c.BaseURL = srv.URL
	albums, err := c.GetArtistAlbums("tok", "xxx")
	if err != nil {
		t.Fatalf("Erreur lors de la récupération des albums : %v", err)
	}
	if len(albums) != 2 {
		t.Fatalf("attendu 2 albums, reçu %d", len(albums))
	}
	if albums[0].Name != "Album 1" {
		t.Errorf("attendu 'Album 1', reçu '%s'", albums[0].Name)
	}
	if albums[1].Name != "Album 2" {
		t.Errorf("attendu 'Album 2', reçu '%s'", albums[1].Name)
	}
}
