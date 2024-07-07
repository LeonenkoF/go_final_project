package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handlers) GetTaskByIdHander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := r.FormValue("id")

	data, err := h.db.GetTaskById(id)
	fmt.Println(data)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Задача не найдена"}`), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Задача не найдена"}`), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
