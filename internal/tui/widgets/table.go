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

func (tw *TableWidget) buildRows() {
	for row, item := range tw.data {
		cellTask := tview.NewTableCell(item["Task"]).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)
		cellCreatedAt := tview.NewTableCell(item["CreatedAt"]).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)
		cellUpdatedAt := tview.NewTableCell(item["UpdatedAt"]).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)
		cellStatus := tview.NewTableCell(item["Status"]).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)

		tw.Table.SetCell(row+1, 0, cellTask)
		tw.Table.SetCell(row+1, 1, cellCreatedAt)
		tw.Table.SetCell(row+1, 2, cellUpdatedAt)
		tw.Table.SetCell(row+1, 3, cellStatus)
	}
}

func (tw *TableWidget) BuildTable() {
	tableHeader := []string{"Task", "CreatedAt", "UpdatedAt", "Status"}

	tw.data = []map[string]string{
		{"Task": "1", "CreatedAt": "1", "UpdatedAt": "1", "Status": "1"},
		{"Task": "2", "CreatedAt": "2", "UpdatedAt": "2", "Status": "2"},
		{"Task": "3", "CreatedAt": "1", "UpdatedAt": "1", "Status": "1"},
		{"Task": "4", "CreatedAt": "1", "UpdatedAt": "2", "Status": "2"},
	}

	tw.Table.SetBorder(true)
	tw.Table.SetSelectable(true, false)

	tw.buildHeader(tableHeader)

	tw.buildRows()

	tw.tableRows = tw.Table.GetRowCount()
}
