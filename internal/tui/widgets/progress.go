package widgets

import (
	"github.com/rivo/tview"
)

type Progress struct {
	ProgressView *tview.Flex
	Gauge        *GaugeView
	commits      *CommitWidget
}

func NewProgress() *Progress {
	p := &Progress{
		tview.NewFlex().SetDirection(tview.FlexRow),
		NewGauge(),
		NewCommitWidget(),
	}
	p.ProgressView.AddItem(p.Gauge, 0, 1, false).
		AddItem(p.commits, 4, 1, false)
	return p
}
