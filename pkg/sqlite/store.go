package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"main/internal/entity"
)

type Store struct {
	db *sql.DB
}

func NewStore(dbFileName string) (*Store, error) {
	const op = "storage.sqlite.New"
	appPath, err := os.Executable()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dbFile := filepath.Join(filepath.Dir(appPath), dbFileName)
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}
	db, err := sql.Open("sqlite", dbFileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if install {
		stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS scheduler(
			id  integer primary key autoincrement,
			date char(8),
			title varchar,
			comment varchar,
			repeat varchar);
		CREATE INDEX IF NOT EXISTS scheduler_date on scheduler(date);
		`)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		_, err = stmt.Exec()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

	}
	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) GetTasks() ([]entity.Task, error) {

	stmt, err := s.db.Query(`SELECT id, date, title, comment, repeat 
	FROM scheduler ORDER 
	BY date ASC
	LIMIT 15;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	data := []entity.Task{}

	for stmt.Next() {
		p := entity.Task{}
		err := stmt.Scan(&p.Id, &p.Date, &p.Title, &p.Comment, &p.Repeat)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		data = append(data, p)
	}
	return data, nil
}

func (s *Store) DeleteTask(Id string) error {

	result, err := s.db.Exec("DELETE FROM scheduler WHERE id=?;", Id)
	if err != nil {
		return err
	}

	value, err := result.RowsAffected()
	if value == 0 {
		return fmt.Errorf("error: %s", err)
	}
	return nil
}

func (s *Store) UpdateTask(input *entity.Task) error {

	result, err := s.db.Exec(`UPDATE scheduler SET 
	date=?, 
	title=?, 
	comment=?, 
	repeat=?
	WHERE id=?;`, input.Date, input.Title, input.Comment, input.Repeat, input.Id)
	if err != nil {
		return err
	}

	value, err := result.RowsAffected()
	if value == 0 {
		return fmt.Errorf("error: %s", err)
	}
	return nil
}

func (s *Store) AddTask(input entity.AddTask) (int64, error) {

	res, err := s.db.Exec(`INSERT INTO scheduler
	(date, title, comment, repeat)
	VALUES(?, ?, ?, ?);`, input.Date, input.Title, input.Comment, input.Repeat)
	if err != nil {
		fmt.Println("error: %s", err)
		return 0, fmt.Errorf("error: %s", err)
	}
	addedId, err := res.LastInsertId()

	if err != nil {
		return addedId, fmt.Errorf("error: %s", err)
	}

	return addedId, nil
}

func (s *Store) GetTaskById(id string) (entity.Task, error) {

	p := entity.Task{}

	stmt := s.db.QueryRow(`SELECT id, date, title, comment, repeat 
	FROM scheduler WHERE id=?;`, id)

	err := stmt.Scan(&p.Id, &p.Date, &p.Title, &p.Comment, &p.Repeat)
	if err != nil {
		return p, err
	}
	return p, nil
}
