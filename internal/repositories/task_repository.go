package repositories

import "database/sql"

type TaskRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db}
}
