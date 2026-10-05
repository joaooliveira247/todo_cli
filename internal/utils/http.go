package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetCommitCount(user string, period time.Time) (*GitHubResponse, error) {
	url := buildCommitSearchURL(user, period)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")

	client := &http.Client{Timeout: 5 * time.Second}

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result *GitHubResponse

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	result.Date = period
	return result, nil
}
