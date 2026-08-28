package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TableWidget struct {
	Table              *tview.Table
	ShowConcludedTasks bool
	tableRows          int
	data               []map[string]string
}

func NewTableWidget() *TableWidget {
	var testMap []map[string]string
	return &TableWidget{tview.NewTable(), false, 0, testMap}
}

func (tw *TableWidget) buildHeader(items []string) {
	if tw.tableRows > 0 {
		tw.Table.Clear()
	}

	for col, item := range items {
		cell := tview.NewTableCell(item).
			SetTextColor(tcell.ColorWhite).
			SetSelectable(false).
			SetAlign(tview.AlignCenter).
			SetExpansion(1)

		tw.Table.SetCell(0, col, cell)
	}
}
