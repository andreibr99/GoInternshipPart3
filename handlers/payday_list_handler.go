package handlers

import (
	"GoInternshipPart3/paydays"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type paydayList struct {
	List []string `json:"dates"`
}

func GetNextPaydays(w http.ResponseWriter, r *http.Request) {
	payDay := strings.TrimPrefix(r.URL.Path, "/till-salary/pay-day/")
	if !strings.HasSuffix(payDay, "/list-dates") {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	payDay = strings.TrimSuffix(payDay, "/list-dates")
	payDayInt, err := strconv.Atoi(payDay)
	if err != nil || payDayInt < 1 || payDayInt > 31 {
		http.Error(w, "Invalid pay day", http.StatusBadRequest)
		return
	}

	dates := paydays.GetPaydaysList(time.Now(), payDayInt)
	resp := paydayList{List: dates}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
