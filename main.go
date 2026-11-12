package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// TODO 1 : host = os.Getenv("HOST"), si vide -> "localhost"
	// TODO 2 : port = os.Getenv("PORT"), si vide -> "8080"
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthcheck)
	// TODO 3 : mux.HandleFunc("GET /health", ...)
	//   dans la fonction : w.WriteHeader(200) puis fmt.Fprint(w, "ok")

	srv := &http.Server{
		Addr:              net.JoinHostPort(host, port), // TODO 4 : host:port (net.JoinHostPort)
		Handler:           mux,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
		// TODO 5 : ReadHeaderTimeout, ReadTimeout, WriteTimeout, IdleTimeout
		//   ex : 5 * time.Second
	}

	log.Printf("écoute sur %s", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	// TODO 6 : err := srv.ListenAndServe() puis log.Fatal(err)
	//_ = fmt.Sprint // supprime quand fmt est utilisé
	//_ = time.Second // supprime quand time est utilisé
	//_ = os.Getenv   // supprime quand os est utilisé
}
func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "ok")
}
