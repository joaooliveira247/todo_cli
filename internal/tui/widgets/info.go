package widgets

import (
	"context"
	"log"
	"time"

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

// TODO: move event loop to app
func (iw *InfoWidget) StartEventLoop() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for now := range ticker.C {
			iw.calendar.UpdateTick(now)
		}
	}()
	return
}

// TODO: DO it with widgets that need default value, can put StartEventLoop here and put it private <- move it to app too
func (iw *InfoWidget) onStartUp() {
	ctx := context.Background()
	if err := iw.commitSync.SyncMissigData(
		ctx,
		iw.calendar.CurrentDate,
		iw.calendar.StartPeriod,
	); err != nil {
		log.Fatal(err)
	}
}

// TODO: check len of commits, do requests and save in db
// INFO: This function only updateDraw, to update values do other, it will
func (iw *InfoWidget) UpdateCommit() error {
	commits, err := iw.commitRepository.GetCommits(iw.calendar.StartPeriod)

	if err != nil {
		return err
	}

	iw.progress.commits.UpdateCommit(commits, iw.calendar.CurrentDate)

	return nil
}
