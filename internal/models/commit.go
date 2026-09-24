package models

import "time"

type CommitModel struct {
	Date        time.Time
	Commits     int
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
