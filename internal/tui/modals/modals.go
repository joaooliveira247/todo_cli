package modals

import (
	"github.com/rivo/tview"
)

type Modals struct {
	pages *tview.Pages
}

func NewModal(pages *tview.Pages) *Modals {
	return &Modals{pages}
}
