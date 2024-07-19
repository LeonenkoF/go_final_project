package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"main/config"
	handler "main/internal/handlers"
	store "main/pkg/sqlite"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Println("Ошибка в создании кофига: %s", err)
		return
	}

	r := chi.NewRouter()

	db, err := store.NewStore("scheduler.db")
	if err != nil {
		log.Println("failed to init storage", err, db)
		return
	}
	h := handler.NewHandler(db)
	h.SetHandlers(r)
	r.Handle("/*", http.FileServer(http.Dir("./web")))

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Println("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
