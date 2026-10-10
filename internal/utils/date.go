package utils

import (
	"time"
)

type DatePeriod struct {
	Start   time.Time
	End     time.Time
	Current time.Time
}

func FormatDate(date time.Time) string {
	return date.Format("02/01/2006")
}

func GetCurrentPeriod() DatePeriod {
	currentDate := time.Now()

	startPerdiod := currentDate.AddDate(0, 0, -int(currentDate.Weekday()))
	endPeriod := currentDate.AddDate(0, 0, (6 - int(currentDate.Weekday())))
	return DatePeriod{
		startPerdiod,
		endPeriod,
		currentDate,
	}
}

func GetCurrentWeekDay() (int, time.Weekday) {
	currentTime := time.Now()

	return int(currentTime.Weekday()), currentTime.Weekday()
}

func IsSameDate(date, target time.Time) bool {
	year, month, day := date.Date()
	yearT, monthT, dayT := target.Date()

	return year == yearT && month == monthT && day == dayT
}

func GetWeekDays() []string {
	return []string{"Sun", "Mon", "Tue", "Wed", "Thi", "Fri", "Sat"}
}
