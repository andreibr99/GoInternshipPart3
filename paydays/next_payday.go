package paydays

import (
	"time"
)

func CalculateNextPayday(today time.Time, payday int) (int, string) {
	paydayTime := time.Date(today.Year(), today.Month(), payday, 0, 0, 0, 0, time.Local)
	lastDayOfMonth := time.Date(today.Year(), today.Month()+1, 0, 0, 0, 0, 0, time.Local).Day()

	if today.Day() == payday {
		if checkForWeekend(today) {
			paydayTime = paydayTime.AddDate(0, 1, 0)
			return int(paydayTime.Sub(today).Hours() / 24), removeWeekends(paydayTime).Format("2006-01-02")
		}
		return 0, paydayTime.Format("2006-01-02")
	}

	// payday is bigger than last day of the month, set payday to last day of the month
	if payday > lastDayOfMonth {
		paydayTime = time.Date(today.Year(), today.Month(), lastDayOfMonth, 0, 0, 0, 0, time.Local)
		paydayTime = removeWeekends(paydayTime)
	} else if today.Equal(paydayTime) || today.After(paydayTime) {
		/*!this else because without it, if the payday is on 31st, and it does not exist in the month, and today
		is the last day of the month, it returns the date and days remaining for the next month payday date.*/

		// payday has already passed this month, calculate next payday for the following month.
		paydayTime = paydayTime.AddDate(0, 1, 0)
		paydayTime = removeWeekends(paydayTime)
	}

	// if today is saturday and the payday is tomorrow, it means that the payday was on friday, so jump to next month
	if checkForWeekend(paydayTime) && paydayTime == today.AddDate(0, 0, 1) {
		paydayTime = paydayTime.AddDate(0, 1, 0)
		paydayTime = removeWeekends(paydayTime)
	}

	paydayTime = removeWeekends(paydayTime)
	daysUntilPayday := int(paydayTime.Sub(today).Hours() / 24)
	dateOfPayday := paydayTime.Format("2006-01-02")

	return daysUntilPayday, dateOfPayday
}

func removeWeekends(payday time.Time) time.Time {
	for checkForWeekend(payday) {
		payday = payday.AddDate(0, 0, -1)
	}
	return payday
}

func checkForWeekend(payday time.Time) bool {
	if payday.Weekday() == time.Saturday || payday.Weekday() == time.Sunday {
		return true
	}
	return false
}
