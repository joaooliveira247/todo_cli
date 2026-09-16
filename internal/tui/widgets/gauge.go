package widgets

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GaugeView struct {
	*tview.Flex
	bar       *gaugeBar
	textLabel *tview.TextView
}

type gaugeBar struct {
	*tview.Box
	percent int
}

func NewGauge() *GaugeView {
	bar := &gaugeBar{
		tview.NewBox(),
		0,
	}

	textLabel := tview.NewTextView().SetTextAlign(tview.AlignCenter)

	gaugeViewArea := tview.NewFlex().SetDirection(tview.FlexRow)

	gaugeViewArea.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexColumn).
			AddItem(tview.NewFlex(), 0, 1, false).
			AddItem(bar, 5, 1, true).
			AddItem(tview.NewFlex(), 0, 1, false),
		0, 1, true).
		AddItem(textLabel, 1, 0, false)
	gaugeViewArea.SetBorder(true).SetTitle(" 📊 Progress ")

	return &GaugeView{
		gaugeViewArea,
		bar,
		textLabel,
	}
}

func (g *gaugeBar) Draw(screen tcell.Screen) {
	g.Box.DrawForSubclass(screen, g)

	x, y, width, height := g.GetInnerRect()
	if width <= 0 || height <= 0 {
		return
	}

	filledHeight := (height * g.percent) / 100

	for row := 0; row < height; row++ {
		currentY := y + height - 1 - row

		style := tcell.StyleDefault.Foreground(tcell.ColorDarkGray).
			Background(tcell.ColorBlack)
		ch := '░'

		if row < filledHeight {
			style = tcell.StyleDefault.Foreground(tcell.ColorGreen).
				Background(tcell.ColorBlack)
			ch = '█'
		}

		for col := 0; col < width; col++ {
			screen.SetContent(x+col, currentY, ch, nil, style)
		}
	}

	text := fmt.Sprintf("%d%%", g.percent)
	textLen := len(text)
	textStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow).
		Background(tcell.ColorBlack).
		Bold(true)

	if textLen <= width {
		centerX := x + (width-textLen)/2
		bottomY := y + height - 1 // Última linha disponível
		for i, char := range text {
			screen.SetContent(centerX+i, bottomY, char, nil, textStyle)
		}
	}
}
