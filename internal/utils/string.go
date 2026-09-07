package utils

import (
	"fmt"

	"github.com/joaooliveira247/todo_cli/internal/models"
)

func FormatStatus(status int) string {
	switch status {
	case models.TaskStatusInProgress:
		return "⏳"
	case models.TaskStatusCannotBeDone:
		return "❌"
	case models.TaskStatusDone:
		return "✅"

	default:
		return "❓"
	}
}

func FormatID(id int) string {
	return fmt.Sprintf("%d", id)
}
