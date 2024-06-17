package main

import (
	"fmt"
	"main/config"
	handler "main/internal/api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.New()

	r := chi.NewRouter()
	handler.SetHandlers(r)

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
