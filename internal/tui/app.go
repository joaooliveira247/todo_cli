package tui

import (
	"database/sql"

	"github.com/gdamore/tcell/v2"
	"github.com/joaooliveira247/todo_cli/internal/repositories"
	"github.com/joaooliveira247/todo_cli/internal/tui/modals"
	"github.com/joaooliveira247/todo_cli/internal/tui/widgets"
	"github.com/rivo/tview"
)

type AppUI struct {
	app        *tview.Application
	pages      *tview.Pages
	modals     *modals.Modals
	table      *widgets.TableWidget
	repository *repositories.TaskRepository
}

func NewAppUI(app *tview.Application, db *sql.DB) *AppUI {
	repository := repositories.NewRepository(db)
	pages := tview.NewPages()
	table := widgets.NewTableWidget()
	modal := modals.NewModal(pages, table)
	table.BuildTable()
	return &AppUI{
		app:        app,
		pages:      pages,
		modals:     modal,
		table:      table,
		repository: repository,
	}
}

func (ui *AppUI) BuildAppUI() *tview.Pages {
	ui.pages.AddPage("main", ui.rootLayout(), true, true)
	ui.app.SetInputCapture(ui.keyPressEvent)

	return ui.pages
}

func (ui *AppUI) keyPressEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyF1:
		ui.modals.ConfirmActionModal(
			"Do you want exit ?",
			"main",
			ui.app.Stop,
		)
		return nil
	}

	return event
}

func (ui *AppUI) contentLayout() *tview.Flex {
	content := tview.NewFlex().
		SetDirection(tview.FlexColumn).AddItem(
		tview.NewBox().SetTitle("Table").SetBorder(true), 0, 1, true,
	).AddItem(tview.NewBox().SetTitle("Progress").SetBorder(true), 40, 0, false)

	content.SetTitle(" 📝 TODO APP ")
	content.SetBorder(true)

	return content
}

func (ui *AppUI) footerLayout() *tview.TextView {
	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[black:yellow] F1 [-:-] Quit")
	return footer
}

func (ui *AppUI) rootLayout() *tview.Flex {
	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ui.contentLayout(), 0, 1, true).
		AddItem(ui.footerLayout(), 1, 0, false)
	return root
}
