package modals

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Modals struct {
	pages *tview.Pages
}

func NewModal(pages *tview.Pages) *Modals {
	return &Modals{pages}
}

func (m *Modals) closeModal(currentModal, backPage string) {
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
