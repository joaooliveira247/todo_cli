package widgets

import (
	"github.com/rivo/tview"
)

type CommitWidget struct {
	*tview.Table
	isLastCheck bool
}
