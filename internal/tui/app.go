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
	footer     *widgets.FooterWidget
}

func NewAppUI(app *tview.Application, db *sql.DB) *AppUI {
	repository := repositories.NewRepository(db)
	pages := tview.NewPages()
	modal := modals.NewModal(pages)
	table := widgets.NewTableWidget(repository, modal)
	footer := widgets.NewFooterWidget()
	return &AppUI{
		app:        app,
		pages:      pages,
		modals:     modal,
		table:      table,
		repository: repository,
		footer:     footer,
	}
}

func (ui *AppUI) BuildAppUI() *tview.Pages {
	ui.pages.AddPage("main", ui.rootLayout(), true, true)
	ui.app.SetInputCapture(ui.keyPressEvent)
	ui.footer.BuildFooter(ui.table.ShowConcludedTasks)
	ui.table.BuildTable()

	return ui.pages
}

func (ui *AppUI) keyPressEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyF1:
		ui.modals.AddTaskModal(ui.repository.InsertTask, ui.table.AddTask)
		return nil
	case tcell.KeyF2:
		ui.table.SetShowConcludedTasks()
		ui.footer.UpdateFooter(ui.table.ShowConcludedTasks)
		return nil
	case tcell.KeyF4:
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
		ui.table.Table, 0, 1, true,
	).AddItem(tview.NewBox().SetTitle("Progress").SetBorder(true), 40, 0, false)

	content.SetTitle(" 📝 TODO APP ")
	content.SetBorder(true)

	return content
}

func (ui *AppUI) rootLayout() *tview.Flex {
	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ui.contentLayout(), 0, 1, true).
		AddItem(ui.footer.Footer, 1, 0, false)
	return root
}
