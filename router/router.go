package router

import (
	"github.com/firas-bouricha/Dev_WEB_API/controller"
	"net/http"
)

// New crée le mux et branche chaque route sur son contrôleur.
func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", controller.Health)
	return mux
}
