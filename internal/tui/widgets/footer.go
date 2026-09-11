package widgets

import (
	"fmt"

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

func (fw *FooterWidget) footerText(isConcludeActive bool) string {
	optionColor := "black:yellow"
	if isConcludeActive {
		optionColor = "white:green"
	}

	return fmt.Sprintf(
		"[black:yellow] F1 [-:-] Add task [%s] F2 [-:-] Show completed tasks [black:yellow] F4 [-:-] Quit",
		optionColor,
	)
}

func (fw *FooterWidget) UpdateFooter(isConcludeActive bool) {
	if fw.Footer != nil {
		fw.Footer.SetText(fw.footerText(isConcludeActive))
	}
}
