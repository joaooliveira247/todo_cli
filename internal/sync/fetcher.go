package sync

import (
	"context"
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

func (ds *DataSync) SyncMissigData(
	ctx context.Context,
	currentDay,
	iniPeriod time.Time,
) error {
	dates, err := ds.repo.GetCommitsIncomplete(iniPeriod, currentDay)

	if err != nil {
		return err
	}

	if err := ds.fecthAndSave(ctx, dates, currentDay); err != nil {
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

	if err := ds.repo.InsertOrUpdateCommit(
		resp.CommitCount,
		currentDay,
		false,
	); err != nil {
		return err
	}

	return nil
}

func (ds *DataSync) fecthAndSave(
	ctx context.Context,
	missingDates []time.Time,
	currentDay time.Time,
) error {
	var responses []*utils.GitHubResponse

	for _, date := range missingDates {
		resp, err := utils.GetCommitCount("joaooliveira247", date)

		if err != nil {
			return err
		}

		responses = append(responses, resp)
	}

	for _, resp := range responses {
		isCompleted := true
		if utils.IsSameDate(resp.Date, currentDay) {
			isCompleted = false
		}
		if err := ds.repo.InsertOrUpdateCommit(
			resp.CommitCount,
			resp.Date,
			isCompleted,
		); err != nil {
			return err
		}
	}

	return nil
}
