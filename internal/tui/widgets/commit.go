package widgets

import (
	"github.com/gdamore/tcell/v2"
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
