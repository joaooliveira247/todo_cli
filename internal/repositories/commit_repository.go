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

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var commits []*models.CommitModel

	for result.Next() {
		commit := &models.CommitModel{}
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

// INFO: Here u can return commits that isn't marked as completed true or in update add logic to safe last out of period
// INFO: maybe only return dates
func (cr *CommitRepository) GetCommitsCompleted(
	iniPeriod time.Time,
) ([]time.Time, error) {
	query := `SELECT date FROM commits WHERE date >= ? AND is_completed = true;`

	rows, err := cr.db.Query(query, iniPeriod.Format("02-01-2006"))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var dates []time.Time

	for rows.Next() {
		var date time.Time

		if err := rows.Scan(&date); err != nil {
			return nil, err
		}

		dates = append(dates, date)
	}

	return dates, nil
}
