package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db := h.db

	id := r.FormValue("id")
	fmt.Println(id)

	err := db.DeleteTask(id)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "ошибка удаления"}`), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{})
}
