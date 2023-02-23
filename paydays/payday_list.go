package paydays

import "time"

func PaydayList(today time.Time, payday int) []string {
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
