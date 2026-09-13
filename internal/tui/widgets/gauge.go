package widgets

import (
	"github.com/rivo/tview"
)

type GaugeView struct {
	*tview.Box
	percent int
}

func NewGauge(app *tview.Application) *GaugeView {
	return &GaugeView{
		tview.NewBox(),
		0,
	}
}

func (g *GaugeView) SetPercent(percent int) *GaugeView {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	g.percent = percent
	return g
}
