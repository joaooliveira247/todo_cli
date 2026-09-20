package repositories

import (
	"database/sql"
)

type CommitRepository struct {
	db *sql.DB
}

func NewCommitRepository(db *sql.DB) *CommitRepository {
	return &CommitRepository{db}
}
