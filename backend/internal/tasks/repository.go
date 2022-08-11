package tasks

import (
	"database/sql"

	. "github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

type Repository interface {
	Get(user_id int64) ([]Task, error)
	GetById(user_id int64, task_id int64) (Task, error)
	Create(task_data *Task) error
	Update(task_data *Task) error
	Delete(user_id int64, task_id int64) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Get(user_id int64) ([]Task, error) {
	q := `
SELECT id, user_id, name, description
FROM tasks
WHERE user_id = ?
	;`

	rows, err := r.db.Query(q, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ts := make([]Task, 0)
	t := Task{}

	for rows.Next() {
		err = rows.Scan(&t.Id, &t.UserId, &t.Name, &t.Description)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	return ts, nil
}

func (r repository) GetById(user_id int64, task_id int64) (Task, error) {
	return Task{}, nil
}

func (r repository) Create(task_data *Task) error {
	return nil
}

func (r repository) Update(task_data *Task) error {
	return nil
}

func (r repository) Delete(user_id int64, task_id int64) error {
	return nil
}