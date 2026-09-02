package models

import (
	"time"

	"github.com/google/uuid"
)

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

func NewTaskModel(task string) *TaskModel {
	return &TaskModel{
		ID:   uuid.NewString(),
		Task: task,
	}
}

func (tm TaskModel) Fields() []string {
	return []string{"ID", "Task", "CreatedAt", "UpdatedAt", "Status"}
}
