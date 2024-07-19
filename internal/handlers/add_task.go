package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"main/internal/entity"
	"main/internal/usecase"
)

func (h *Handlers) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	if len(input.Date) == 0 || input.Date == "" {
		input.Date = time.Now().Format("20060102")

	}

	_, err = time.Parse("20060102", input.Date)
	if input.Date != " " {
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "дата представлена в формате, отличном от 20060102"}`), http.StatusBadRequest)
			return
		}
	}
	if input.Date < time.Now().Format("20060102") {
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
	addedId, err := db.AddTask(input)
	insertId := strconv.FormatInt(addedId, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusBadRequest)
		return
	}

	res := entity.Task{Id: insertId}
	respBytes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusBadRequest)
		return
	}

	_, err = w.Write(respBytes)
	if err != nil {
		log.Println("Ошибка при записи данных в ResponseWriter")
	}
}
