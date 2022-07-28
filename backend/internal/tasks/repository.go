package tasks

import (
	"database/sql"
	"log"
)

type Task interface {}  // TODO: implement task structure
type Tasks []Task

type Repository interface {
	Get(user_id int64) (Tasks, error)
	GetById(user_id int64, task_id int64) (Task, error)
	Create(user_id int64, task_data Task) error
	Update(user_id int64, task_id int64, task_data Task) error
	Delete(user_id int64, task_id int64) error
}

type repository struct {
	db *sql.DB
	logger *log.Logger
}

func (r *repository) Get(user_id int64) (Tasks, error) {
	return nil, nil
}

func (r *repository) GetById(user_id int64, task_id int64) (Task, error) {
	return nil, nil
}

func (r *repository) Create(user_id int64, task_data Task) error {
	return nil
}

func (r *repository) Update(user_id int64, task_id int64, task_data Task) error {
	return nil
}

func (r *repository) Delete(user_id int64, task_id int64) error {
	return nil
}