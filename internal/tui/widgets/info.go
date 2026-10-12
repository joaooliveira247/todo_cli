package widgets

import (
	"github.com/joaooliveira247/todo_cli/internal/repositories"
	"github.com/joaooliveira247/todo_cli/internal/sync"
	"github.com/rivo/tview"
)

// INFO: InfoWidget carry calendar(), progress(Gauge and commit)
type InfoWidget struct {
	Info             *tview.Flex
	app              *tview.Application
	taskRepository   *repositories.TaskRepository
	commitRepository *repositories.CommitRepository
	commitSync       *sync.DataSync
	calendar         *Calendar
	progress         *Progress
}
