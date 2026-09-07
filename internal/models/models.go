package models

import (
	"time"

	"github.com/rivo/tview"
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

type TaskRow struct {
	ID        *tview.TableCell
	Task      *tview.TableCell
	CreatedAt *tview.TableCell
	UpdatedAt *tview.TableCell
	Status    *tview.TableCell
}

func NewTaskModel(task string) *TaskModel {
	return &TaskModel{
		Task: task,
	}
}

func (tm TaskModel) Fields() []string {
	return []string{"ID", "Task", "CreatedAt", "UpdatedAt", "Status"}
}
