package cmd

import (
	"github.com/joaooliveira247/todo_cli/internal/tui"
	"github.com/rivo/tview"
)

type App struct {
	tviewApp *tview.Application
}

func (app *App) Run() error {
	if err := app.tviewApp.Run(); err != nil {
		return err
	}
	return nil
}

func NewApp() *App {
	app := &App{tviewApp: tview.NewApplication()}
	ui := tui.NewAppUI()
	app.tviewApp.SetRoot(ui.BuildAppUI(), true)
	app.tviewApp.SetInputCapture(app.keyPressEvent)

	return app
}
