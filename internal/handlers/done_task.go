package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"main/internal/usecase"
)

func (h *Handlers) DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db := h.db

	id := r.FormValue("id")

	if id == "" {
		http.Error(w, fmt.Sprintf(`{"error": "id не заполнен"}`), http.StatusBadRequest)
		return
	}

	task, err := db.GetTaskById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "задача не найдена"}`), http.StatusBadRequest)
		return
	}

	if task.Repeat == "" || len(task.Repeat) == 0 {
		db.DeleteTask(id)
	} else {

		now := time.Now()

		nextDate, err := usecase.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "ошибка вычисления следующей даты"}`), http.StatusBadRequest)
			return
		}

		task.Date = nextDate

		err = db.UpdateTask(&task)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "задача не обновлена"}`), http.StatusBadRequest)
			return
		}

	}
	json.NewEncoder(w).Encode(map[string]interface{}{})
}
