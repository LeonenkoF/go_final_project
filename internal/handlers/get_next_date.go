package handler

import (
	"fmt"
	"log"
	"main/internal/usecase"
	"net/http"
	"time"
)

func GetNextDate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	prsednow, _ := time.Parse("20060102", now)
	fmt.Println(repeat)

	nextdate, err := usecase.NextDate(prsednow, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Ошибка при выполнении функции GetNextDate: %s", err.Error())
		return
	}
	w.Write([]byte(nextdate))

}
