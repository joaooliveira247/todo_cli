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

func (tr *TaskRepository) InsertTask(task *models.TaskModel) error {
	query := `INSERT INTO list (task) VALUES (?);`

	tx, err := tr.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err := tx.Exec(query, task.Task); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
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
