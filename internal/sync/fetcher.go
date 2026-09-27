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
