package sqlite

import (
	"fmt"
	"main/internal/entity"
)

func (s *DBManager) GetTasks() ([]entity.Task, error) {
	const op = "storage.sqlite.GetAll"

	stmt, err := s.db.Query(getTasksQuery)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()
	data := []entity.Task{}

	for stmt.Next() {
		p := entity.Task{}
		err := stmt.Scan(&p.Id, &p.Date, &p.Title, &p.Comment, &p.Repeat)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		data = append(data, p)
	}
	return data, nil
}

func (s *DBManager) DeleteTask(id int) error {
	const op = "storage.sqlite.DeleteTask"

	_, err := s.db.Exec(deleteTaskQuery, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *DBManager) UpdateTask(input entity.Task) error {
	const op = "storage.sqlite.UpdateTask"

	_, err := s.db.Exec(updateTaskQuery, input)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *DBManager) AddTask(input entity.Task) error {
	const op = "storage.sqlite.AddTask"

	_, err := s.db.Exec(addTaskQuery, input)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *DBManager) getTaskById(id int) error {
	const op = "storage.sqlite.AddTask"

	_, err := s.db.Exec(getTaskByIdQuery, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
