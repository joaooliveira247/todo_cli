package utils

import "time"

func FormatDate(date time.Time) string {
	return date.Format("02/01/2006")
}

func GetCurrentPeriod() (time.Time, time.Time) {
	currentDate := time.Now()

	startPerdiod := currentDate.AddDate(0, 0, -int(currentDate.Weekday()))
	endPeriod := currentDate.AddDate(0, 0, (6 - int(currentDate.Weekday())))
	return startPerdiod, endPeriod
}
