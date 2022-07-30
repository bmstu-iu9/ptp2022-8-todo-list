package tasks

import (
	. "github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

type CreateTaskRequest struct {
	Name 		string 	`json:"name"`
	Description *string `json:"description"`
}

type UpdateTaskRequest struct {
	Name 		*string `json:"name"`
	Description *string `json:"description"`
}

type Service interface {
	Get(user_id int64) ([]Task, error)
	GetById(user_id int64, task_id int64) (Task, error)
	Create(user_id int64, task_data CreateTaskRequest) error
	Update(user_id int64, task_id int64, task_data UpdateTaskRequest) error
	Delete(user_id int64, task_id int64) error
}

type service struct {
	r *repository
}

func (s *service) Get(user_id int64) ([]Task, error) {
	tasks, err := s.r.Get(user_id)

	if err != nil {
		return nil, err
	}

	ret := make([]Task, len(tasks))
	for i, t := range tasks {
		ret[i] = Task{
			Id: t.Id,
			Name: t.Name,
			Description: t.Description,
		}
	}

	return ret, err
}

func (s *service) GetById(user_id int64, task_id int64) (Task, error) {
	task, err := s.r.GetById(user_id, task_id)
	
	if err != nil {
		return Task{}, err
	}

	ret := Task{
		Id: task.Id,
		Name: task.Name,
		Description: task.Description,
	}

	return ret, err
}

func (t *CreateTaskRequest) Validate() error {
	// TODO: implement validation of creation
	return nil
}

func (s *service) Create(user_id int64, task_data CreateTaskRequest) error {
	err := task_data.Validate()

	if err != nil {
		return err
	}

	task := &Task{
		UserId: user_id,
		Name: task_data.Name,
		Description: task_data.Description,
	}

	err = s.r.Create(task)

	return err
}

func (t *UpdateTaskRequest) Validate() error {
	// TODO: implement update validation
	return nil
}

func (s *service) Update(user_id int64, task_id int64, task_data UpdateTaskRequest) error {
	err := task_data.Validate()

	if err != nil {
		return err
	}

	task, err := s.r.GetById(user_id, task_id)

	if err != nil {
		return err
	}

	var u bool = false

	if task_data.Name != nil {
		u = true
		task.Name = *task_data.Name
	}

	if task_data.Description != nil {
		u = true
		task.Description = task_data.Description
	}

	if u {
		err = s.r.Update(&task)
	}

	return err
}

func (s *service) Delete(user_id int64, task_id int64) error {
	err := s.r.Delete(user_id, task_id)
	return err
}