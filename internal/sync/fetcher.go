package sync

import (
	"context"
	"log"
	"time"

	"github.com/joaooliveira247/todo_cli/internal/repositories"
	"github.com/joaooliveira247/todo_cli/internal/utils"
)

// TODO: DataSync only run in startup app, put sync in app
type DataSync struct {
	repo *repositories.CommitRepository
}

func NewDataSync(repo *repositories.CommitRepository) *DataSync {
	return &DataSync{repo}
}

// INFO: this func 'll run in event loop, app will call from database
// INFO: Call it next day in loop event
// INFO: check if one or more past days is in report
func (ds *DataSync) SyncMissigData(
	ctx context.Context,
	currentDay,
	iniPeriod time.Time,
) error {
	// separe and check what dates are present, when start will check iniPeriod - 1
	iniPeriod = iniPeriod.AddDate(0, 0, -1)

	dates, err := ds.repo.GetCommitsCompleted(iniPeriod)

	if err != nil {
		return err
	}

	// TODO: create a func to update and create today that will check if exist too, sync only sync past
	if !utils.ContainsSameDate(dates, currentDay) {
		if err := ds.repo.InsertCommitCount(0, currentDay); err != nil {
			log.Fatal("contais")
			return err
		}
		dates = append(dates, currentDay)
	}

	//INFO: date already formated in utils.http
	if err := ds.fecthAndSave(ctx, dates); err != nil {
		log.Fatal("fetch")
		return err
	}

	return nil
}

func (ds *DataSync) fecthAndSave(
	ctx context.Context,
	missingDates []time.Time,
) error {
	// INFO: for now only return nil, after test it
	var responses []*utils.GitHubResponse

	for _, date := range missingDates {
		resp, err := utils.GetCommitCount("joaooliveira247", date)

		if err != nil {
			return err
		}

		responses = append(responses, resp)
	}

	for _, resp := range responses {
		if err := ds.repo.UpdateCommitCount(
			resp.Date,
			resp.CommitCount,
			true,
		); err != nil {
			log.Fatal("fetch UpdateCommit", err)
			return err
		}
	}

	return nil
}
