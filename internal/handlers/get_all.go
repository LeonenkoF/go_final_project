package handler

import (
	"encoding/json"
	"fmt"
	"main/internal/entity"
	"net/http"
)

func (h *Handlers) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	db := h.db

	if r.Method == "GET" {

		data, err := db.GetTasks()
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusBadRequest)
			fmt.Printf("err")
			return
		}

		m := make(map[string][]entity.Task)
		m["tasks"] = data

		if m["tasks"] == nil {
			m["tasks"] = []entity.Task{}
		}

		resp, err := json.Marshal(m)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusBadRequest)
			fmt.Printf("err")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

	}
}
