package repositories

import "database/sql"

type TaskRepository struct {
	db *sql.DB
}
