package widgets

import (
	"fmt"
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

func (c *Calendar) createClock() *Calendar {
	c.clock = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(fmt.Sprintf("[yellow]%s[-]", c.CurrentDate.Format("15:04:05")))
	return c
}

func (c *Calendar) createDate() *Calendar {
	currentDate := c.CurrentDate.Format("Monday 02/01/2006")
	c.date = tview.NewTextView().
		SetDynamicColors(true).
		SetText(fmt.Sprintf("[yellow]%s[-]", currentDate)).
		SetTextAlign(tview.AlignCenter)

	return c
}

func (c *Calendar) createPeriod() *Calendar {
	c.periodView = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	c.periodView.SetText(
		fmt.Sprintf(
			"🚀 %s - 🏁 %s",
			utils.FormatDate(c.StartPeriod),
			utils.FormatDate(c.EndPeriod),
		),
	)
	c.periodView.SetBorder(true).SetTitle(" ⌛ Period ")
	return c
}

func (c *Calendar) buildCalendar() {
	c.createClock().createDate().createPeriod()

	clockView := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(c.date, 1, 0, false).
		AddItem(c.clock, 1, 0, false)
	clockView.SetBorder(true).SetTitle(" 📆 Calendar ")

	c.CalendarView.AddItem(clockView, 4, 0, false).
		AddItem(c.periodView, 3, 0, false)
}
