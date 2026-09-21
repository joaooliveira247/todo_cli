package repositories

import (
	"database/sql"
	"time"

	"github.com/joaooliveira247/todo_cli/internal/models"
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

func (cr *CommitRepository) GetCommits(
	iniPeriod time.Time,
) ([]*models.CommitModel, error) {
	result, err := cr.db.Query(
		`SELECT * FROM commits WHERE date >= ?;`,
		iniPeriod,
	)
	defer result.Close()

	if err != nil {
		return nil, err
	}

	var commits []*models.CommitModel

	for result.Next() {
		var commit *models.CommitModel
		if err := result.Scan(
			&commit.Date,
			&commit.Commits,
			&commit.IsCompleted,
			&commit.CreatedAt,
			&commit.UpdatedAt,
		); err != nil {
			return nil, err
		}
		commits = append(commits, commit)
	}

	return commits, nil
}
