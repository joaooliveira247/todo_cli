package cmd

import (
	"github.com/rivo/tview"
)

type AppUI struct {
	pages *tview.Pages
}

func NewAppUI() *AppUI {
	return &AppUI{pages: tview.NewPages()}
}

func contentLayout() *tview.Flex {
	content := tview.NewFlex().
		SetDirection(tview.FlexColumn).AddItem(
		tview.NewBox().SetTitle("Table").SetBorder(true), 0, 1, true,
	).AddItem(tview.NewBox().SetTitle("Progress").SetBorder(true), 40, 0, false)

	content.SetTitle(" 📝 TODO APP ")
	content.SetBorder(true)

	return content
}

func footerLayout() *tview.TextView {
	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[black:yellow] F1 [-:-] Quit")
	return footer
}

func RootLayout() *tview.Flex {
	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(contentLayout(), 0, 1, true).
		AddItem(footerLayout(), 1, 0, false)
	return root
}
