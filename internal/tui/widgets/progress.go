package widgets

import (
	"github.com/rivo/tview"
)

type Progress struct {
	ProgressView *tview.Flex
	Gauge        *GaugeView
	commits      *CommitWidget
}
