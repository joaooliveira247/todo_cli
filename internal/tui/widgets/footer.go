package widgets

import (
	"github.com/rivo/tview"
)

type FooterWidget struct {
	Footer *tview.TextView
}

func NewFooterWidget() *FooterWidget {
	return &FooterWidget{
		tview.NewTextView().
			SetDynamicColors(true).
			SetTextAlign(tview.AlignCenter),
	}
}
