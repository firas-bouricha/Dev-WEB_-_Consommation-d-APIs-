package main

import (
	"github.com/firas-bouricha/Dev_WEB_API/controller"
	"github.com/firas-bouricha/Dev_WEB_API/router"
	"github.com/firas-bouricha/Dev_WEB_API/spotify"

	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	} else {
		return fallback
	}
}
func main() {
	//avant go run . set -a; source .env; set +a pour prendre host et port dans le fichier .env
	host := getenv("HOST", "localhost")
	port := getenv("PORT", "8080")
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("SPOTIFY_CLIENT_ID et SPOTIFY_CLIENT_SECRET sont obligatoires")
	}
	// on crée un client Spotify avec un timeout de 10 secondes
	c := spotify.NewClient(10 * time.Second)
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	//on recupère le token pour l'utiliser dans les appels à l'API Spotify
	tok, err := c.FetchToken()
	if err != nil {
		log.Fatalf("Impossible de récupérer le token : %v", err)
	}
	// on récupère tous les albums de l'artiste Jul
	albums, err := c.GetArtistAlbums(tok.AccessToken, spotify.JulID)
	// en cas d'erreur on log et on quitte le programme
	if err != nil {
		log.Fatalf("Impossible de récupérer les albums : %v", err)
	}
	// si ca marche on log le nombre d'albums récupérés
	log.Printf("Albums de Jul : %d", len(albums))
	// comme albulm ici on recupère tous les titres de tous les albums de Jul meme avec les doublons
	tracks, err := c.GetAllTracks(tok.AccessToken, albums)
	if err != nil {
		log.Fatalf("Impossible de récupérer les titres : %v", err)
	}
	// on log le nombre de titres récupérés
	log.Printf("Titres de Jul : %d", len(tracks))
	// on supprime les doublons de la liste de titres et on log le nombre de titres sans doublons
	controller.Tracks = spotify.SansDoublons(tracks)
	log.Printf("Titres de Jul  sans doublons : %d", len(controller.Tracks))

	// Serveur HTTP
	srv := &http.Server{
		Addr:              net.JoinHostPort(host, port),
		ReadHeaderTimeout: 5 * time.Second,  // max pour envoyer l'en-tête
		ReadTimeout:       15 * time.Second, // max pour lire toute la requête
		WriteTimeout:      15 * time.Second, // max pour écrire la réponse
		IdleTimeout:       60 * time.Second, // ferme une connexion inactive
		Handler:           router.New(),     // associe chaque URL (/health, /...) à la fonction qui répond
	}

	log.Printf("écoute sur %s", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

}
