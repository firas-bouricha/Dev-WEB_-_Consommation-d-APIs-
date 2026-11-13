package controller

import (
	"html/template"
	"net/http"

	"github.com/firas-bouricha/Dev_WEB_API/spotify"
)

// PageData regroupe ce que le template affiche.
type PageData struct {
	Tracks []spotify.Track
}

// Home affiche la liste des titres.
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		http.Error(w, "Erreur interne", 500)
		return
	}
	data := PageData{Tracks: Tracks}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur interne", 500)
	}
}
