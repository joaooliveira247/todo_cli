package cmd

import "github.com/rivo/tview"

func footerLayout() *tview.TextView {
	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[black:yellow] F1 [-:-] Quit")
	return footer
}

func RootLayout() *tview.Flex {
	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Content"), 0, 1, true).
		AddItem(tview.NewTextView().SetText("Footer").SetTextAlign(tview.AlignCenter), 1, 0, false)
	return root
}
