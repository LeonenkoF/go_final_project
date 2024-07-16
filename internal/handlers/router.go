package handler

import (
	store "main/pkg/sqlite"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	db *store.Store
}

func NewHandler(db *store.Store) *Handlers {
	return &Handlers{db: db}
}

func (h *Handlers) SetHandlers(router chi.Router) {
	router.Post("/api/task", h.AddTaskHandler)
	router.Post("/api/task/done", h.DoneTaskHandler)
	router.Get("/api/tasks", h.GetAllHandler)
	router.Get("/api/nextdate", GetNextDate)
	router.Get("/api/task", h.GetTaskByIdHander)
	router.Put("/api/task", h.UpdateTaskHandler)
	router.Delete("/api/task", h.DeleteTaskHandler)
}
