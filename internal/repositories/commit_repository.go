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

// TODO: update here to on clofict do nothing
func (cr *CommitRepository) InsertCommitCount(
	commits int,
	date time.Time,
) error {
	query := `INSERT INTO commits (commits, date) VALUES (?, ?) ON CONFLICT (date) DO NOTHING;`

	tx, err := cr.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err := tx.Exec(
		query,
		commits,
		date.Format("2006-01-02"),
	); err != nil {
		return err
	}

	return tx.Commit()
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

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err := tx.Exec(query, commits, isCompleted, date); err != nil {
		return err
	}

	return tx.Commit()
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
// TODO: change name of this func
func (cr *CommitRepository) GetCommitsIncomplete(
	iniPeriod,
	currentDay time.Time,
) ([]time.Time, error) {
	query := `WITH RECURSIVE period(date) AS (
  SELECT datetime(?, '-1 days')
  UNION ALL
  SELECT datetime(date, '+1 days')
  FROM period
  WHERE date <= ?
)
SELECT p.date
FROM period p
LEFT JOIN commits c ON p.date = c.date
WHERE c.date IS NULL OR c.is_completed = false;`

	rows, err := cr.db.Query(
		query,
		iniPeriod.Format("2006-01-02"),
		currentDay.Format("2006-01-02"),
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var dates []time.Time

	for rows.Next() {
		var dateString string

		if err := rows.Scan(&dateString); err != nil {
			return nil, err
		}

		parseDate, err := time.Parse("2006-01-02", dateString[:10])

		if err != nil {
			return nil, err
		}

		dates = append(dates, parseDate)
	}

	return dates, nil
}
