package modals

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Modals struct {
	pages *tview.Pages
}

func NewModal(pages *tview.Pages) *Modals {
	return &Modals{pages}
}

func (m *Modals) closeModal(currentModal, backPage string, closeDelay int) {
	time.Sleep(time.Second * time.Duration(closeDelay))
	m.pages.RemovePage(currentModal)
	m.pages.SwitchToPage(backPage)
}

func (m *Modals) modalNavigation(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyRight:
		return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone)
	case tcell.KeyLeft:
		return tcell.NewEventKey(tcell.KeyBacktab, 0, tcell.ModNone)
	}
	return event
}

func (m *Modals) ConfirmActionModal(msg, backModal string, doneFunc func()) {
	modal := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"Yes", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "Yes":
				doneFunc()
			case "Cancel":
				m.closeModal("confirmActionModal", "main")
			}
		})
	modal.SetInputCapture(m.modalNavigation)

	m.pages.AddPage("confirmActionModal", modal, true, true)
}
