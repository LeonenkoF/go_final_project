package handler

import (
	"main/pkg/sqlite"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var webDir = "./web"

type Handlers struct {
	db *sqlite.DBManager
}

func NewHandler(db *sqlite.DBManager) *Handlers {
	return &Handlers{db: db}
}

func (h *Handlers) SetHandlers(router chi.Router) {
	FileServer(router, "/", http.Dir(webDir))
	router.Post("/api/task", h.AddTaskHandler)
	router.Post("/api/task/done", h.DoneTaskHander)
	router.Get("/api/tasks", h.GetAllHandler)
	router.Get("/api/task", h.GetTaskByIdHander)
	router.Put("/api/task", h.UpdateTaskHandler)
}
