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

// GetNextPaydays handles an HTTP request and returns a JSON response with a list of dates for the months left until end
// of the year, given a pay day parameter in the URL. The pay day parameter is expected to be an integer between
// 1 and 31 and is extracted from the URL path. If the URL is invalid or does not end with "/list-dates", an
// HTTP 400 error is returned. If the pay day parameter is invalid, an HTTP 400 error is returned. The function
// uses the GetPaydaysList function from the paydays package to calculate the dates. The response is returned in JSON
// format with a list of dates. If there is an error marshaling the response to JSON, the function returns an HTTP 500
// error.
func GetNextPaydays(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
