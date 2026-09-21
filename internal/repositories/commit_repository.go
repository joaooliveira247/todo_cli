package repositories

import (
	"database/sql"
	"time"
)

type CommitRepository struct {
	db *sql.DB
}

func NewCommitRepository(db *sql.DB) *CommitRepository {
	return &CommitRepository{db}
}

func (cr *CommitRepository) InsertCommitCount(
	commits int,
	date time.Time,
) error {
	query := `INSERT INTO commits (commits, date) VALUES ?, ?;`

	tx, err := cr.db.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(query, commits, date); err != nil {
		return err
	}

	return nil
}
