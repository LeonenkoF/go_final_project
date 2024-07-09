package sqlite

import (
	"fmt"
	"main/internal/entity"

	_ "modernc.org/sqlite"
)

func (s *DBManager) GetTasks() ([]entity.Task, error) {

	stmt, err := s.db.Query(getTasksQuery)
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

func (s *DBManager) DeleteTask(Id string) error {

	result, err := s.db.Exec(deleteTaskQuery, Id)
	if err != nil {
		return err
	}

	value, err := result.RowsAffected()
	if value == 0 {
		return fmt.Errorf("error")
	}
	return nil
}

func (s *DBManager) UpdateTask(input *entity.Task) error {

	result, err := s.db.Exec(updateTaskQuery, input.Date, input.Title, input.Comment, input.Repeat, input.Id)
	if err != nil {
		return err
	}

	value, err := result.RowsAffected()
	if value == 0 {
		return fmt.Errorf("error")
	}
	return nil
}

func (s *DBManager) AddTask(input entity.AddTask) int64 {

	res, err := s.db.Exec(addTaskQuery, input.Date, input.Title, input.Comment, input.Repeat)
	if err != nil {
	}
	addedId, _ := res.LastInsertId()

	return addedId
}

func (s *DBManager) GetTaskById(id string) (entity.Task, error) {

	p := entity.Task{}

	stmt := s.db.QueryRow(getTaskByIdQuery, id)

	err := stmt.Scan(&p.Id, &p.Date, &p.Title, &p.Comment, &p.Repeat)
	if err != nil {
		return p, err
	}
	return p, nil
}
