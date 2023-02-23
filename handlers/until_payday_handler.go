package handlers

import (
	"GoInternshipPart3/paydays"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type paydayDate struct {
	Days int    `json:"days"`
	Date string `json:"date"`
}

func HowMuchUntilPayday(w http.ResponseWriter, r *http.Request) {
	payday, err := strconv.Atoi(r.URL.Query().Get("pay_day"))
	if err != nil || payday < 1 || payday > 31 {
		http.Error(w, "Invalid pay_day query parameter", http.StatusBadRequest)
		return
	}

	days, date := paydays.CalculateNextPayday(time.Now(), payday)
	resp := paydayDate{Days: days, Date: date}

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
