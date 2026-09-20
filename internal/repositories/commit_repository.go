package repositories

import (
	"database/sql"
)

type CommitRepository struct {
	db *sql.DB
}
