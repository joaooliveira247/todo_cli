package models

import "time"

const (
	TaskStatusInProgress = iota
	TaskStatusDone
	TaskStatusCannotBeDone
)

type TaskModel struct {
	ID        int
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    int
}
