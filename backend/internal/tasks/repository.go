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

func (r *repository) Get(user_id int64) ([]Task, error) {
	return nil, nil
}

func (r *repository) GetById(user_id int64, task_id int64) (Task, error) {
	return Task{}, nil
}

func (r *repository) Create(task_data *Task) error {
	return nil
}

func (r *repository) Update(task_data *Task) error {
	return nil
}

func (r *repository) Delete(user_id int64, task_id int64) error {
	return nil
}