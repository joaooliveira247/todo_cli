package widgets

import (
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
