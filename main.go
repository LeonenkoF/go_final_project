package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"main/config"
	handler "main/internal/handlers"
	store "main/pkg/sqlite"
)

func main() {

	cfg := config.New()

	r := chi.NewRouter()

	db, err := store.NewStore("scheduler.db")
	if err != nil {
		log.Fatal("failed to init storage", err, db)
		os.Exit(1)
	}
	h := handler.NewHandler(db)
	h.SetHandlers(r)
	r.Handle("/*", http.FileServer(http.Dir("./web")))

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
