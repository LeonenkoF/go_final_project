package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var webDir = "./web"

func SetHandlers(router chi.Router) {
	FileServer(router, "/", http.Dir(webDir))
	router.Get("/api/nextdate", GetNextDate)
}
