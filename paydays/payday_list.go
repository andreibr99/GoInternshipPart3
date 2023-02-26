package paydays

import "time"

// GetPaydaysList takes as input a type time variable which represents the desired day to be in, and an integer
// which represents the pay day. It uses the function CalculateNextPayday to get the next pay day date for each month,
// until the end of the year. It returns a slice of strings, with the next pay day dates.
func GetPaydaysList(today time.Time, payday int) []string {
	var dates []string
	months := 12
	// prevent taking into account 1st month of next year
	if payday < today.Day() {
		months = 11
	}

	for m := today.Month(); m <= time.Month(months); m++ {
		_, date := CalculateNextPayday(today, payday)

		dates = append(dates, date)
		today = today.AddDate(0, 1, 0)
	}
	return dates
}
