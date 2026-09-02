package modals

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/joaooliveira247/todo_cli/internal/tui/widgets"
	"github.com/rivo/tview"
)

type LogType = tcell.Color

const (
	LogLevelSuccess = LogType(tcell.ColorLightGreen)
	LogLevelError   = LogType(tcell.ColorRed)
)

type Modals struct {
	pages *tview.Pages
	table *widgets.TableWidget
}

func NewModal(
	pages *tview.Pages,
	table *widgets.TableWidget,
) *Modals {
	return &Modals{pages, table}
}

func (m *Modals) closeModal(currentModal, backPage string, closeDelay int) {
	time.Sleep(time.Second * time.Duration(closeDelay))
	m.pages.RemovePage(currentModal)
	m.pages.SwitchToPage(backPage)
}

func (m *Modals) modalNavigation(
	currentModal, backPage string,
) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRight:
			return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone)
		case tcell.KeyLeft:
			return tcell.NewEventKey(tcell.KeyBacktab, 0, tcell.ModNone)
		case tcell.KeyESC:
			m.closeModal(currentModal, backPage, 0)
			return nil
		}
		return event
	}
}

func (m *Modals) customModal(
	item tview.Primitive,
	width, height int,
) *tview.Flex {
	modal := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(item, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)

	return modal
}

func (m *Modals) AddTaskModal() {
	var taskValue string
	modalName := "addTaskModal"

	form := tview.NewForm().
		AddInputField("Task", "", 30, nil, func(text string) { taskValue = text }).
		AddButton("Save", func() {
			if taskValue == "" {
				m.LogMessageModal(
					"Field \"Task\" can't be empty",
					LogLevelError,
				)
				return
			}
			// logic to safe task
			m.LogMessageModal("Task Added", LogLevelSuccess)
			return
		}).AddButton("Cancel", func() {
		m.closeModal(modalName, "main", 0)
	}).SetButtonsAlign(tview.AlignCenter)
	form.SetBorder(true)

	modal := m.customModal(form, 40, 7)

	modal.SetInputCapture(m.modalNavigation(modalName, "main"))
	m.pages.AddPage(modalName, modal, true, true)
}

func (m *Modals) ConfirmActionModal(msg, backModal string, doneFunc func()) {
	modal := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"Yes", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "Yes":
				doneFunc()
			case "Cancel":
				m.closeModal("confirmActionModal", "main", 0)
			}
		})
	modal.SetInputCapture(m.modalNavigation("confirmActionModal", "main"))

	m.pages.AddPage("confirmActionModal", modal, true, true)
}

func (m *Modals) LogMessageModal(msg string, level LogType) {
	var modalName string

	switch level {
	case LogLevelError:
		modalName = "LogLevelError"
	case LogLevelSuccess:
		modalName = "LogLevelSuccess"
	}

	//TODO: check width and height of msg, current max is 35 chars
	textBox := tview.NewTextView().
		SetText(fmt.Sprintf("\n\n[black]%s\n\n", msg)).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	textBox.SetBackgroundColor(level).SetBorder(true)

	modal := m.customModal(textBox, 35, 7)

	modal.SetInputCapture(m.modalNavigation(modalName, "main"))
	m.pages.AddPage(modalName, modal, true, true)
}
