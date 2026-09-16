package repositories

import (
	"database/sql"
	"time"

	"github.com/joaooliveira247/todo_cli/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (tr *TaskRepository) InsertTask(
	task string,
) (*models.TaskModel, error) {
	query := `INSERT INTO list (task) VALUES (?) RETURNING *;`

	tx, err := tr.db.Begin()

	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var newTask models.TaskModel

	if err := tx.QueryRow(query, task).
		Scan(&newTask.ID, &newTask.Task, &newTask.CreatedAt, &newTask.UpdatedAt, &newTask.Status); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &newTask, nil
}

func (tr *TaskRepository) UpdateTask(task *models.TaskModel) error {
	query := `UPDATE list
	SET task = ?, updated_at = ?, status = ?
	WHERE id = ?
	RETURNING task, updated_at, status;`

	tx, err := tr.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if err := tx.QueryRow(query, task.Task, time.Now(), task.Status, task.ID).
		Scan(&task.Task, &task.UpdatedAt, &task.Status); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) GetTasks(
	iniPeriod time.Time,
	completed bool,
) ([]*models.TaskModel, error) {
	// returns tasks that is not completed if completed is true return all completed in that period
	// by default always return only tasks in progress
	var tasks []*models.TaskModel
	var args []any

	query := `SELECT * FROM list WHERE status = 0 ORDER BY created_at ASC;`

	if completed {
		query = `SELECT * FROM list WHERE created_at >= ? OR updated_at >= ? OR status = 0 ORDER BY created_at ASC;`
		args = append(args, iniPeriod, iniPeriod)
	}

	rows, err := tr.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		task := &models.TaskModel{}

		err := rows.Scan(
			&task.ID,
			&task.Task,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.Status,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *TaskRepository) GetStats(
	period time.Time,
) (*models.TaskStats, error) {
	var stats models.TaskStats
	query := `WITH cte_stats AS (
	SELECT
		COUNT(status) AS total,
		COUNT(CASE WHEN status = 1 THEN 1 END) AS completed
	FROM list
	WHERE created_at >= ? OR updated_at >= ?)
	SELECT
	total,
	completed,
	COALESCE(ROUND((completed * 100.0) / NULLIF(total, 0), 0), 0) AS percentage
	FROM cte_stats;`

	if err := tr.db.QueryRow(query, period, period).
		Scan(&stats.Total, &stats.Completed, &stats.Percentage); err != nil {
		return nil, err
	}

	return &stats, nil
}
