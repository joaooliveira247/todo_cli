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

func NewInfoWidget(
	app *tview.Application,
	taskReposiotry *repositories.TaskRepository,
	commitRepository *repositories.CommitRepository,
) *InfoWidget {
	//WARN: widgets that need base from database values need be build in father widget
	sync := sync.NewDataSync(commitRepository)
	info := &InfoWidget{
		tview.NewFlex().SetDirection(tview.FlexRow),
		app,
		taskReposiotry,
		commitRepository,
		sync,
		NewCalendar(app),
		NewProgress(),
	}
	return info
}
