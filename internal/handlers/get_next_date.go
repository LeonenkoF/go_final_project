package handler

import (
	"fmt"
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

	nextdate, err := usecase.NextDate(prsednow, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("error GetNextDate func: %s", err)
		return
	}
	w.Write([]byte(nextdate))

}
