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

func InitDB() (*sql.DB, error) {
	dbPath, err := getDatabasePath()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	return db, nil
}
