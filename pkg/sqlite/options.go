package repository

import (
	"fmt"
	"main/internal/entity"

	_ "modernc.org/sqlite"
)

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
		return fmt.Errorf("error")
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
		return fmt.Errorf("error")
	}
	return nil
}

func (s *Store) AddTask(input entity.AddTask) int64 {

	res, err := s.db.Exec(`INSERT INTO scheduler
	(date, title, comment, repeat)
	VALUES(?, ?, ?, ?);`, input.Date, input.Title, input.Comment, input.Repeat)
	if err != nil {
	}
	addedId, _ := res.LastInsertId()

	return addedId
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
