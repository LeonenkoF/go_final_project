package main

import (
	"log"
	"main/config"
	handler "main/internal/handlers"
	repository "main/pkg/sqlite"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {

	cfg := config.New()

	r := chi.NewRouter()

	db, err := repository.NewStore("scheduler.db")
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
