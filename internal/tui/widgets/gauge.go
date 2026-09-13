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
