package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"main/internal/usecase"
)

func GetNextDate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	prsednow, err := time.Parse("20060102", now)
	if err != nil {
		http.Error(w, fmt.Sprintln(`{"error":"ошибка формата поля "now""}`), http.StatusBadRequest)
		return
	}

	nextdate, err := usecase.NextDate(prsednow, date, repeat)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err), http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(nextdate))
	if err != nil {
		log.Println("Ошибка при записи данных в ResponseWriter")
		return
	}

}
