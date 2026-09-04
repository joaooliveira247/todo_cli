package widgets

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/joaooliveira247/todo_cli/internal/models"
	"github.com/joaooliveira247/todo_cli/internal/repositories"
	"github.com/joaooliveira247/todo_cli/internal/utils"
	"github.com/rivo/tview"
)

type TableWidget struct {
	Table              *tview.Table
	ShowConcludedTasks bool
	tableRows          int
	Data               []*models.TaskModel
	repository         *repositories.TaskRepository
}

func NewTableWidget(repository *repositories.TaskRepository) *TableWidget {
	//TODO: fix iniPeriod, and error handling here
	data, _ := repository.GetTasks(time.Now(), false)
	return &TableWidget{tview.NewTable(), false, 0, data, repository}
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
		cellID := tview.NewTableCell(utils.FormatID(item.ID)).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)
		cellTask := tview.NewTableCell(item.Task).
			SetMaxWidth(40).
			SetAlign(tview.AlignCenter)
		cellCreatedAt := tview.NewTableCell(utils.FormatDate(item.CreatedAt)).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)
		cellUpdatedAt := tview.NewTableCell(utils.FormatDate(item.UpdatedAt)).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)
		cellStatus := tview.NewTableCell(utils.FormatStatus(item.Status)).
			SetExpansion(1).
			SetAlign(tview.AlignCenter)

		tw.Table.SetCell(row+1, 0, cellID)
		tw.Table.SetCell(row+1, 1, cellTask)
		tw.Table.SetCell(row+1, 2, cellCreatedAt)
		tw.Table.SetCell(row+1, 3, cellUpdatedAt)
		tw.Table.SetCell(row+1, 4, cellStatus)
	}
}

func (tw *TableWidget) BuildTable() {
	tw.Table.SetBorder(true)
	tw.Table.SetSelectable(true, false)

	tw.buildHeader()

	tw.buildRows()

	tw.tableRows = tw.Table.GetRowCount()
}
