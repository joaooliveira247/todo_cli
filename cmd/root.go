package cmd

import (
	"github.com/rivo/tview"
)

type App struct {
	tviewApp *tview.Application
}

func NewApp() *App {
	app := &App{tviewApp: tview.NewApplication()}

	return app
}
