package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/entity"
	"main/internal/usecase"
	"net/http"
	"strconv"
	"time"
)

func (h *Handlers) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

	db := h.db

	var input entity.Task

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Задача не найдена"}`), http.StatusBadRequest)
		return
	}

	_, err = strconv.Atoi(input.Id)

	if input.Id == "" || len(input.Id) == 0 || err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "не указан id"}`), http.StatusBadRequest)
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

	err = db.UpdateTask(&input)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Не удалось отредактировать задачу"}`), http.StatusInternalServerError)
		return
	}
	res := entity.Task{}
	respBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Ошибка при преобразовании в JSON формат: %s", err.Error())
		return
	}
	w.Write(respBytes)
}
