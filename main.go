package main

import (
	"log"
	"main/config"
	handler "main/internal/handlers"
	"main/pkg/sqlite"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.New()

	r := chi.NewRouter()
	handler.SetHandlers(r)

	db, err := sqlite.New()
	if err != nil {
		log.Fatal("failed to init storage", err, db)
		os.Exit(1)
	}

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
