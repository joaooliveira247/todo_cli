package widgets

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/joaooliveira247/todo_cli/internal/models"
	"github.com/rivo/tview"
)

type CommitWidget struct {
	*tview.Table
	isLastCheck bool
}

func NewCommitWidget() *CommitWidget {
	c := &CommitWidget{tview.NewTable(), false}
	c.createCommitsArea()
	return c
}

// In case change day don't forget to back cell black or every cell 'll be green
func (cw *CommitWidget) UpdateCommit(
	commits []*models.CommitModel,
	currentDay time.Time,
) {
	for col, commit := range commits {
		if col == int(currentDay.Weekday()) {
			cw.Table.GetCell(0, col).SetBackgroundColor(tcell.ColorGreenYellow)
		}
		cw.Table.GetCell(1, col).SetText(fmt.Sprint(commit.Commits))
	}
}

func (cw *CommitWidget) createCommitsArea() *CommitWidget {
	fields := []string{"Sun", "Mon", "Tue", "Wed", "Thi", "Fri", "Sat"}
	if cw.Table.GetRowCount() > 0 {
		cw.Table.Clear()
	}

	for col, item := range fields {
		headerCell := tview.NewTableCell(item).
			SetTextColor(tcell.ColorWhite).
			SetBackgroundColor(tcell.ColorBlack).
			SetSelectable(false).
			SetAlign(tview.AlignCenter).
			SetExpansion(1)
		contentCell := tview.NewTableCell("0").
			SetTextColor(tcell.ColorWhite).
			SetBackgroundColor(tcell.ColorBlack).
			SetSelectable(false).
			SetAlign(tview.AlignCenter).
			SetExpansion(1)

		cw.Table.SetCell(0, col, headerCell)
		cw.Table.SetCell(1, col, contentCell)
	}

	cw.Table.SetBorder(true)
	cw.Table.SetTitle("  Commits ")

	return cw
}
