package utils

import (
	"fmt"
	"time"
)

type GitHubResponse struct {
	Date        time.Time
	CommitCount int `json:"total_count"`
}

func buildCommitSearchURL(author string, t time.Time) string {
	localTime := t.Local()

	startOfDay := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0,
		0,
		0,
		0,
		localTime.Location(),
	)
	endOfDay := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		23,
		59,
		59,
		0,
		localTime.Location(),
	)

	isoLayout := "2006-01-02T15:04:05-07:00"

	dateRange := fmt.Sprintf(
		"%s..%s",
		startOfDay.Format(isoLayout),
		endOfDay.Format(isoLayout),
	)

	query := fmt.Sprintf("author:%s+author-date:%s", author, dateRange)

	return fmt.Sprintf("https://api.github.com/search/commits?q=%s", query)
}
