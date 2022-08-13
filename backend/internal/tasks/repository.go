package tasks

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

type Repository interface {
	Get(user_id int64) ([]entity.Task, error)
	GetById(task_id int64) (entity.Task, error)
	Create(task_data *entity.Task) error
	Update(task_data *entity.Task) error
	Delete(task_id int64) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Get(user_id int64) ([]entity.Task, error) {
	q := "SELECT id, user_id, name, description FROM tasks WHERE user_id = $1;"

	rows, err := r.db.Query(q, user_id)

	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	tasks := make([]entity.Task, 0)
	task := entity.Task{}

	for rows.Next() {
		err = rows.Scan(&task.Id, &task.UserId, &task.Name, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r repository) GetById(task_id int64) (entity.Task, error) {
	q := "SELECT id, user_id, name, description FROM tasks WHERE id = $1"

	t := entity.Task{}
	err := r.db.QueryRow(q, task_id).Scan(&t.Id, &t.UserId, &t.Name, &t.Description)

	return t, err
}

func (r repository) Create(task_data *entity.Task) error {
	q := "INSERT INTO tasks(user_id, name, description) VALUES ($1, $2, $3) RETURNING id;"
	err := r.db.QueryRow(q, task_data.UserId, task_data.Name, task_data.Description).Scan(&task_data.Id)
	return err
}

func (r repository) Update(task_data *entity.Task) error {
	q := "UPDATE tasks SET (name, description) = ($1, $2) WHERE id = $3;"
	_, err := r.db.Exec(q, task_data.Name, task_data.Description, task_data.Id)
	return err
}

func (r repository) Delete(task_id int64) error {
	q := "DELETE FROM tasks WHERE id = $1;"
	_, err := r.db.Exec(q, task_id)
	return err
}