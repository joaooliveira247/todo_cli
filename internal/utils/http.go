package utils

import (
	"time"
)

type GitHubResponse struct {
	Date        time.Time
	CommitCount int `json:"total_count"`
}
