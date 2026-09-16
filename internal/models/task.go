package models

import (
	"time"

	"github.com/joaooliveira247/todo_cli/internal/utils"
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

type TaskStats struct {
	Total      int
	Completed  int
	Percentage int
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

func (tm *TaskModel) ToRow() *TaskRow {
	cellID := tview.NewTableCell(utils.FormatID(tm.ID)).
		SetExpansion(1).
		SetAlign(tview.AlignCenter).SetReference(tm)
	cellTask := tview.NewTableCell(tm.Task).
		SetMaxWidth(40).
		SetAlign(tview.AlignCenter)
	cellCreatedAt := tview.NewTableCell(utils.FormatDate(tm.CreatedAt)).
		SetExpansion(1).
		SetAlign(tview.AlignCenter)
	cellUpdatedAt := tview.NewTableCell(utils.FormatDate(tm.UpdatedAt)).
		SetExpansion(1).
		SetAlign(tview.AlignCenter)
	cellStatus := tview.NewTableCell(utils.FormatStatus(tm.Status)).
		SetExpansion(1).
		SetAlign(tview.AlignCenter)

	return &TaskRow{cellID, cellTask, cellCreatedAt, cellUpdatedAt, cellStatus}
}

func (tm TaskModel) Fields() []string {
	return []string{"ID", "Task", "CreatedAt", "UpdatedAt", "Status"}
}
