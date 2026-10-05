package utils

import (
	"fmt"
)

func FormatStatus(status int) string {
	switch status {
	case 0:
		return "⏳"
	case 2:
		return "❌"
	case 1:
		return "✅"

	default:
		return "❓"
	}
}

func ParseDropDownOption(status string) int {
	switch status {
	case "⏳ InProgress":
		return 0
	case "❌ CannotBeDone":
		return 2
	case "✅ Done":
		return 1

	default:
		return -1
	}
}

func FormatID(id int) string {
	return fmt.Sprintf("%d", id)
}
