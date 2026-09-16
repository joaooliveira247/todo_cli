package widgets

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type gaugeBar struct {
	*tview.Box
	percent int
}

func NewGauge(app *tview.Application) *gaugeBar {
	return &gaugeBar{
		tview.NewBox(),
		0,
	}
}

func (g *gaugeBar) SetPercent(percent int) *gaugeBar {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	g.percent = percent
	return g
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
