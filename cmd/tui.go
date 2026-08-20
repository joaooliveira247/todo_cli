package cmd

import "github.com/rivo/tview"

func RootLayout() *tview.Flex {
	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Content"), 0, 1, true).
		AddItem(tview.NewTextView().SetText("Footer").SetTextAlign(tview.AlignCenter), 1, 0, false)
	return root
}
