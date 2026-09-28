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
	dates, err := ds.repo.GetCommitsIncomplete(iniPeriod, currentDay)

	if err != nil {
		return err
	}

	if err := ds.fecthAndSave(ctx, dates); err != nil {
		return err
	}

	return nil
}

func (ds *DataSync) EventUpdate(
	ctx context.Context,
	currentDay time.Time,
) error {
	resp, err := utils.GetCommitCount("joaooliveira247", currentDay)

	if err != nil {
		return err
	}

	if err := ds.repo.UpdateCommitCount(
		currentDay,
		resp.CommitCount,
		false,
	); err != nil {
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
