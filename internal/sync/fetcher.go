package sync

import (
	"github.com/joaooliveira247/todo_cli/internal/repositories"
)

// TODO: DataSync only run in startup app, put sync in app
type DataSync struct {
	repo *repositories.CommitRepository
}
