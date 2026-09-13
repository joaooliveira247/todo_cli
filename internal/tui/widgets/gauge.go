package widgets

import (
	"github.com/rivo/tview"
)

type GaugeView struct {
	*tview.Box
	percent int
}
