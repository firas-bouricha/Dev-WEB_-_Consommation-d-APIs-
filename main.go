package main

import (
	"github.com/firas-bouricha/Dev_WEB_API/router"
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
	_ = clientID
	_ = clientSecret

	srv := &http.Server{
		Addr:              net.JoinHostPort(host, port),
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
		Handler:           router.New(),
	}

	log.Printf("écoute sur %s", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

}
