package handler

import (
	repository "main/pkg/sqlite"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	db *repository.Store
}

func NewHandler(db *repository.Store) *Handlers {
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
