package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joaooliveira247/todo_cli/internal/models"
	"github.com/rivo/tview"
)

type TableWidget struct {
	Table              *tview.Table
	ShowConcludedTasks bool
	tableRows          int
	Data               []map[string]string
}

func NewTableWidget() *TableWidget {
	var testMap []map[string]string
	return &TableWidget{tview.NewTable(), false, 0, testMap}
}

func (tw *TableWidget) buildHeader() {
	var model models.TaskModel
	if tw.tableRows > 0 {
		tw.Table.Clear()
	}

	for col, item := range model.Fields() {
		cell := tview.NewTableCell(item).
			SetTextColor(tcell.ColorWhite).
			SetSelectable(false).
			SetAlign(tview.AlignCenter).
			SetExpansion(1)

		tw.Table.SetCell(0, col, cell)
	}
}

func (tw *TableWidget) buildRows() {
	for row, item := range tw.Data {
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
	tw.Table.SetBorder(true)
	tw.Table.SetSelectable(true, false)

	tw.buildHeader()

	tw.buildRows()

	tw.tableRows = tw.Table.GetRowCount()
}
