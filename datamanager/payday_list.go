package datamanager

import "time"

func PaydayList(payday int) []string {
	today := time.Now()
	endOfYear := time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Local)
	paydayDate := time.Date(time.Now().Year(), time.Now().Month(), payday, 0, 0, 0, 0, time.Local)
	if paydayDate.Before(today) || paydayDate.Equal(today) {
		paydayDate = paydayDate.AddDate(0, 1, 0)
	}
	var dates []string
	for paydayDate.Before(endOfYear) {
		dates = append(dates, paydayDate.Format("2006-01-02"))
		paydayDate = paydayDate.AddDate(0, 1, 0)
	}
	return dates
}
