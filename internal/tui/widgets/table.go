package widgets

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/joaooliveira247/todo_cli/internal/models"
	"github.com/joaooliveira247/todo_cli/internal/repositories"
	"github.com/joaooliveira247/todo_cli/internal/tui/modals"
	"github.com/rivo/tview"
)

type TableWidget struct {
	Table              *tview.Table
	ShowConcludedTasks bool
	tableRows          int
	Data               []*models.TaskModel
	repository         *repositories.TaskRepository
	modal              *modals.Modals
}

func NewTableWidget(repository *repositories.TaskRepository, modals *modals.Modals) *TableWidget {
	//TODO: fix iniPeriod, and error handling here
	data, _ := repository.GetTasks(time.Now(), false)
	return &TableWidget{tview.NewTable(), false, 0, data, repository, modals}
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
	for rowIdx, item := range tw.Data {
		row := item.ToRow()

		tw.Table.SetCell(rowIdx+1, 0, row.ID)
		tw.Table.SetCell(rowIdx+1, 1, row.Task)
		tw.Table.SetCell(rowIdx+1, 2, row.CreatedAt)
		tw.Table.SetCell(rowIdx+1, 3, row.UpdatedAt)
		tw.Table.SetCell(rowIdx+1, 4, row.Status)
	}
}

func (tw *TableWidget) BuildTable() {
	tw.Table.SetBorder(true)
	tw.Table.SetSelectable(true, false)

	tw.buildHeader()

	tw.buildRows()

	tw.tableRows = tw.Table.GetRowCount()
}
