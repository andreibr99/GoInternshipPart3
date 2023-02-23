package main

import (
	"GoInternshipPart3/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/till-salary/how-much", handlers.HowMuchUntilPayday)
	http.HandleFunc("/till-salary/pay-day/", handlers.GetNextPaydays)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
