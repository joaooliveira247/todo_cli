package widgets

import (
	"time"

	"github.com/joaooliveira247/todo_cli/internal/utils"
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

func NewCalendar(app *tview.Application) *Calendar {
	startPeriod, EndPeriod := utils.GetCurrentPeriod()
	c := &Calendar{
		tview.NewFlex().SetDirection(tview.FlexRow),
		app,
		nil,
		nil,
		nil,
		startPeriod,
		EndPeriod,
		time.Now(),
	}
	c.buildCalendar()
	return c
}
