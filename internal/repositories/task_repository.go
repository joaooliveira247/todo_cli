package repositories

import (
	"database/sql"

	"github.com/joaooliveira247/todo_cli/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (tr *TaskRepository) InsertTask(
	task *models.TaskModel,
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

	if err := tx.QueryRow(query, task.Task).
		Scan(&newTask.ID, &newTask.Task, &newTask.CreatedAt, &newTask.UpdatedAt, &newTask.UpdatedAt, &newTask.Status); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &newTask, nil
}

func (tr *TaskRepository) ChangeTaskStatus(id int, status int) error {
	query := `UPDATE list SET status = ? WHERE id = ?;`

	tx, err := tr.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err := tx.Exec(query, id, status); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
