package widgets

import (
	"time"

	"github.com/rivo/tview"
)

type Calendar struct {
	CalendarView *tview.Flex
	app          *tview.Application
	clock        *tview.TextView
	date         *tview.TextView
	periodView   *tview.TextView

	StartPeriod time.Time
	EndPeriod   time.Time
	CurrentDate time.Time
}
