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

// HowMuchUntilPayday handles an HTTP request and returns a JSON response with information about how many days until
// the next payday, given a pay day parameter in the query string. The pay day parameter must be a number between
// 1 and 31. The function uses the CalculateNextPayday function from the paydays package to calculate the number of
// days until the next payday and the date of the next payday. The response is returned in JSON format with the number
// of days and the date. If there is an error parsing the pay day parameter, the function returns an HTTP 400 error.
// If there is an error marshaling the response to JSON, the function returns an HTTP 500 error.
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
