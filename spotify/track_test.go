package spotify

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAlbumTracks(t *testing.T) {
	// faux serveur : une page de 2 titres, pas de page suivante
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"items":[{"id":"t1","name":"Titre 1"},{"id":"t2","name":"Titre 2"}],"next":null,"total":2}`)
	}))
	defer srv.Close()

	c := NewClient(2 * time.Second)
	c.BaseURL = srv.URL

	// album d'exemple : ses infos doivent être recopiées dans chaque titre
	album := Album{ID: "a1", Name: "Mon album", ReleaseDate: "2019-06-28", AlbumType: "album"}

	tracks, err := c.GetAlbumTracks("tok", album)
	if err != nil {
		t.Fatalf("erreur : %v", err)
	}
	if len(tracks) != 2 {
		t.Fatalf("attendu 2 titres, reçu %d", len(tracks))
	}
	if tracks[0].Album.Name != "Mon album" {
		t.Errorf("album non recopié : %q", tracks[0].Album.Name)
	}
	if tracks[1].Album.ReleaseDate != "2019-06-28" {
		t.Errorf("date non recopiée : %q", tracks[1].Album.ReleaseDate)
	}
}

func TestGetAllTracks(t *testing.T) {
	// même faux serveur : chaque album renvoie 2 titres
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"items":[{"id":"t1","name":"Titre 1"},{"id":"t2","name":"Titre 2"}],"next":null,"total":2}`)
	}))
	defer srv.Close()

	c := NewClient(2 * time.Second)
	c.BaseURL = srv.URL

	albums := []Album{{ID: "a1", Name: "A1"}, {ID: "a2", Name: "A2"}}

	tracks, err := c.GetAllTracks("tok", albums)
	if err != nil {
		t.Fatalf("erreur : %v", err)
	}
	// 2 albums x 2 titres = 4
	if len(tracks) != 4 {
		t.Fatalf("attendu 4 titres, reçu %d", len(tracks))
	}
	// les 2 derniers titres viennent du 2e album
	if tracks[3].Album.Name != "A2" {
		t.Errorf("mauvais album : %q", tracks[3].Album.Name)
	}
}
