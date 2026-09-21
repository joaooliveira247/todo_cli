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

func (cr *CommitRepository) UpdateCommitCount(
	date time.Time,
	commits int,
	isCompleted bool,
) error {
	query := `UPDATE commits SET commits = ?, is_completed = ? WHERE date = ?;`

	tx, err := cr.db.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(query, commits, isCompleted, date); err != nil {
		return err
	}

	return nil
}
