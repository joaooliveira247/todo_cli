package models

import "time"

type TaskModel struct {
	ID        int
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    int
}
