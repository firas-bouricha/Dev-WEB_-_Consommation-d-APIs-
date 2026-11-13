package router

import (
	"github.com/firas-bouricha/Dev_WEB_API/controller"
	"net/http"
)

// Le ServeMux (mux) : renvoie l’annuaire avec toutes les routes
func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", controller.Health)
	mux.HandleFunc("GET /{$}", controller.Home)
	return mux
}
