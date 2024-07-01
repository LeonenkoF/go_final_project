package handler

import (
	"encoding/json"
	"fmt"
	"main/internal/entity"
	"main/internal/usecase"
	"net/http"
	"strconv"
	"time"
)

func (h *Handlers) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		db := h.db

		var input entity.AddTask

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "ошибка десериализации JSON"}`), http.StatusBadRequest)
			return
		}

		if len(input.Title) < 1 {
			http.Error(w, fmt.Sprintf(`{"error": "не указан заголовок задачи"}`), http.StatusBadRequest)
			return
		}

		parsedDate, err := time.Parse("20060102", input.Date)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "дата представлена в формате, отличном от 20060102"}`), http.StatusBadRequest)
			return
		}

		if len(input.Date) == 0 || input.Date == "" {
			input.Date = time.Now().Format("20060102")

		} else if parsedDate.Before(time.Now()) {
			if len(input.Repeat) == 0 || input.Repeat == "" {
				input.Date = time.Now().Format("20060102")

			} else {
				input.Date, err = usecase.NextDate(time.Now(), input.Date, input.Repeat)
				if err != nil {
					http.Error(w, fmt.Sprintf(`{"error": "правило повторения указано в неправильном формате"}`), http.StatusBadRequest)
					return
				}
			}
		}
		insertId := strconv.FormatInt(db.AddTask(input), 10)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusBadRequest)
			return
		}

		response := entity.Task{Id: insertId}
		responseBytes, _ := json.Marshal(response)
		w.Write(responseBytes)
	}
}
