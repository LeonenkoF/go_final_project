package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DBManager struct {
	db *sql.DB
}

const dbFileName = "scheduler.db"

func New() (*DBManager, error) {
	const op = "storage.sqlite.New"
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dbFile := filepath.Join(filepath.Dir(appPath), dbFileName)
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}
	db, err := sql.Open("sqlite", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	if install {
		stmt, err := db.Prepare(createTableQuery)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		_, err = stmt.Exec()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

	}
	return &DBManager{db: db}, nil
}

func (s *DBManager) Close() error {
	return s.db.Close()
}
