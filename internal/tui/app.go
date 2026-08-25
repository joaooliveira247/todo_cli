package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joaooliveira247/todo_cli/internal/tui/modals"
	"github.com/rivo/tview"
)

type AppUI struct {
	app    *tview.Application
	pages  *tview.Pages
	modals *modals.Modals
}

func NewAppUI(app *tview.Application) *AppUI {
	return &AppUI{app: app, pages: tview.NewPages()}
}

func (ui *AppUI) BuildAppUI() *tview.Pages {
	ui.pages.AddPage("main", ui.rootLayout(), true, true)
	ui.app.SetInputCapture(ui.keyPressEvent)

	return ui.pages
}

func (ui *AppUI) keyPressEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyF1:
		ui.ConfirmActionModalLayout(
			"Do you want exit ?",
			"main",
			ui.app.Stop,
		)
		return nil
	}

	return event
}

func (ui *AppUI) ConfirmActionModalLayout(
	msg, backPage string,
	doneFunc func(),
) {
	modal := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"Yes", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "Yes":
				doneFunc()
			case "Cancel":
				ui.pages.RemovePage("confirm_page")
				ui.pages.SwitchToPage(backPage)
			}
		})
	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRight:
			return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone)
		case tcell.KeyLeft:
			return tcell.NewEventKey(tcell.KeyBacktab, 0, tcell.ModNone)
		}
		return event
	})

	ui.pages.AddPage("confirm_modal", modal, true, true)
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
