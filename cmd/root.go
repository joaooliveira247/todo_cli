package cmd

import (
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

	return app
}
