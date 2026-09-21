package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func getDatabasePath() (string, error) {
	// only for linux, maybe later for others OS
	dataDir := os.Getenv("XDG_DATA_HOME")

	if dataDir == "" {
		home, err := os.UserHomeDir()

		if err != nil {
			return "", err
		}
		dataDir = filepath.Join(home, ".local", "share")
	}

	appDir := filepath.Join(dataDir, "todo-cli")

	if err := os.Mkdir(appDir, 0755); err != nil && !os.IsExist(err) {
		return "", err
	}

	return filepath.Join(appDir, "todo.db"), nil
}

func createTable(db *sql.DB) error {
	tx, err := db.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS list (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		status INT DEFAULT 0
		);`,
	); err != nil {
		return err
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS commits (
		date DATETIME PRIMARY KEY,
		commits INTEGER DEFAULT 0,
		is_completed BOOLEAN DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func InitDB() (*sql.DB, error) {
	dbPath, err := getDatabasePath()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		return nil, err
	}

	err = createTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}
