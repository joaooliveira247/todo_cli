package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	tviewApp *tview.Application
	ui       *AppUI
}

func (app *App) keyPressEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyF1:
		app.tviewApp.Stop()
		return nil
	}

	return event
}

func (app *App) Run() error {
	if err := app.tviewApp.Run(); err != nil {
		return err
	}
	return nil
}

func NewApp() *App {
	ui := NewAppUI()
	app := &App{tviewApp: tview.NewApplication()}
	app.tviewApp.SetRoot(ui.BuildAppUI(), true)
	app.tviewApp.SetInputCapture(app.keyPressEvent)

	return app
}
