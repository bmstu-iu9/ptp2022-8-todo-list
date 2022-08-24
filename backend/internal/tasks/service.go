package tasks

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

type CreateTaskRequest struct {
	UserId					int64				`json:"-"`
	Name 					string 				`json:"name"`
	Description 			string 				`json:"description,omitempty"`
	CreatedOn				string				`json:"createdOn"`
	DueDate					string				`json:"dueDate"`
	SchtirlichHumorescue 	string 				`json:"schtirlichHumorescue"`
	Labels					[]entity.TaskLabel	`json:"labels"`
	Status					string				`json:"status"`
}

type UpdateTaskRequest struct {
	TaskId		int64				`json:"-"`
	Name 		string 				`json:"name,omitempty"`
	Description string 				`json:"description,omitempty"`
	DueDate		string				`json:"dueDate,omitempty"`
	Labels		[]entity.TaskLabel	`json:"labels,omitempty"`
	Status		string				`json:"status,omitempty"`
}

type Service interface {
	Get(user_id int64) ([]entity.Task, error)
	GetById(task_id int64) (entity.Task, error)
	Create(task_data *CreateTaskRequest) (entity.Task, error)
	Update(task_data *UpdateTaskRequest) (entity.Task, error)
	Delete(task_id int64) (entity.Task, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) Get(user_id int64) ([]entity.Task, error) {
	tasks, err := s.r.Get(user_id)

	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (s service) GetById(task_id int64) (entity.Task, error) {
	task, err := s.r.GetById(task_id)
	
	if err != nil {
		return entity.Task{}, err
	}

	return task, err
}

func (t *CreateTaskRequest) Validate() error {
	// TODO: implement validation of creation
	return nil
}

func (s service) Create(task_data *CreateTaskRequest) (entity.Task, error) {
	err := task_data.Validate()

	if err != nil {
		return entity.Task{}, err
	}

	task := &entity.Task{
		UserId: task_data.UserId,
		Name: task_data.Name,
		Description: task_data.Description,
		CreatedOn: task_data.CreatedOn,
		DueDate: task_data.DueDate,
		SchtirlichHumorescue: task_data.SchtirlichHumorescue,
		Labels: task_data.Labels,
		Status: task_data.Status,
	}

	// logger := log.New()
	// logger.Debug("Labels proccessed: ", task.Labels)

	err = s.r.Create(task)

	return *task, err
}

func (t *UpdateTaskRequest) Validate() error {
	// TODO: implement update validation
	return nil
}

func (s service) Update(task_data *UpdateTaskRequest) (entity.Task, error) {
	err := task_data.Validate()

	if err != nil {
		return entity.Task{}, err
	}

	task, err := s.r.GetById(task_data.TaskId)

	if err != nil {
		return entity.Task{}, err
	}

	or := func (ss ...string) (string) { for _, s := range ss { if s != "" { return s } }; return ""; }

	task.Name = or(task_data.Name, task.Name)
	task.Description = or(task_data.Description, task.Description)
	task.DueDate = or(task_data.DueDate, task.DueDate)
	task.Labels = task_data.Labels
	task.Status = or(task_data.Status, task.Status)

	err = s.r.Update(&task)

	return task, err
}

func (s service) Delete(task_id int64) (entity.Task, error) {
	task, err := s.r.GetById(task_id)
	if err != nil {
		return entity.Task{}, err
	}

	err = s.r.Delete(task_id)

	return task, err
}