package datamanager

import (
	"time"
)

func DaysUntilNextPayday(payday int) (int, string) {
	if payday < 1 || payday > 31 {
		return -1, ""
	}
	today := time.Now()
	paydayTime := time.Date(time.Now().Year(), time.Now().Month(), payday, 0, 0, 0, 0, time.Local)
	if today.Equal(paydayTime) || today.After(paydayTime) {
		paydayTime = paydayTime.AddDate(0, 1, 0)
	}
	daysUntilPayday := int(paydayTime.Sub(today).Hours()/24) + 1
	dateOfPayday := paydayTime.Format("2006-01-02")
	return daysUntilPayday, dateOfPayday
}
