package tasks

import (
	"fmt"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/validation"
)

// CreateTaskRequest represents task creation request
// description is optional
type CreateTaskRequest struct {
	UserId               int64  `json:"-"`
	Name                 Name   `json:"name"`
	Description          Text   `json:"description,omitempty"`
	CreatedOn            Date   `json:"createdOn"`
	DueDate              Date   `json:"dueDate"`
	SchtirlichHumorescue Text   `json:"schtirlichHumorescue"`
	Labels               Labels `json:"labels"`
	Status               Status `json:"status"`
}

// UpdateTaskRequest represents task  modify request
// all of the fields is optional
type UpdateTaskRequest struct {
	TaskId               int64  `json:"-"`
	Name                 Name   `json:"name,omitempty"`
	Description          Text   `json:"description,omitempty"`
	DueDate              Date   `json:"dueDate,omitempty"`
	SchtirlichHumorescue Text   `json:"schtirlichHumorescue,omitempty"`
	Labels               Labels `json:"labels,omitempty"`
	Status               Status `json:"status,omitempty"`
}

type (
	Name   entity.Name
	Text   entity.Text
	Date   entity.Date
	Labels []Label
	Status entity.Status
	Color  entity.Color
	Label  struct {
		Id     int64 `json:"id"`
		TaskId int64 `json:"-"`
		Name   Name  `json:"text"`
		Color  Color `json:"color"`
	}
)

func (f *Name) validate() bool {
	if f == nil {
		return false
	}
	return validation.ValidateField(string(*f), 1, 255, ".*")
}

func (f *Text) validate() bool {
	if f == nil {
		return false
	}
	return validation.ValidateField(string(*f), 1, 8192, ".*")
}

func (f *Date) validate() bool {
	if f == nil {
		return false
	}

	return validation.ValidateField(string(*f), 1, 100, `^(19|20)\d\d-(0[1-9]|1[012])-(0[1-9]|[12]\d|3[01])T[0-5]\d:[0-5]\d:[0-5]\dZ$`)
}

func (f *Labels) validate() bool {
	if f == nil {
		return false
	}

	for _, label := range *f {
		if !label.validate() {
			return false
		}
	}
	return true
}

func (f *Color) validate() bool {
	if f == nil {
		return false
	}

	return validation.ValidateField(string(*f), 0, 255, `^#[0-9A-Fa-f]{6}$`)
}

func (f *Label) validate() bool {
	if f == nil {
		return false
	}

	if f.Id != 0 {
		return true
	}

	return f.Name.validate() && f.Color.validate()
}

func (f *Status) validate() bool {
	if f == nil {
		return false
	}

	return validation.ValidateField(string(*f), 0, 255, "^(in progress|done|outdated)$")
}

// Service encapsulates usecase logic for tasks.
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

// NewService creates a new user service.
func NewService(r Repository) Service {
	return service{r}
}

// Get returns all tasks for user with specified id
func (s service) Get(user_id int64) ([]entity.Task, error) {
	tasks, err := s.r.Get(user_id)

	if err != nil {
		return nil, err
	}

	return tasks, err
}

// GetById returns single task with specified id
func (s service) GetById(task_id int64) (entity.Task, error) {
	task, err := s.r.GetById(task_id)

	if err != nil {
		return entity.Task{}, err
	}

	return task, err
}

func (t *CreateTaskRequest) Validate() error {
	if !(t.Name.validate() &&
		t.Description.validate() &&
		t.CreatedOn.validate() &&
		t.DueDate.validate() &&
		t.SchtirlichHumorescue.validate() &&
		t.Labels.validate() &&
		t.Status.validate()) {
		return errors.ErrValidation
	}
	return nil
}

func toEntityLabels(labels Labels) entity.Labels {
	lbs := entity.Labels{}
	for _, lb := range labels {
		lbs = append(lbs, entity.TaskLabel{
			Id:     lb.Id,
			TaskId: lb.TaskId,
			Name:   entity.Name(lb.Name),
			Color:  entity.Color(lb.Color),
		})
	}
	return lbs
}

// Create creates task from task_data argument
func (s service) Create(task_data *CreateTaskRequest) (entity.Task, error) {
	err := task_data.Validate()

	if err != nil {
		return entity.Task{}, fmt.Errorf("%w: %v", errors.ErrValidation, err)
	}

	task := &entity.Task{
		UserId:               task_data.UserId,
		Name:                 entity.Name(task_data.Name),
		Description:          entity.Text(task_data.Description),
		CreatedOn:            entity.Date(task_data.CreatedOn),
		DueDate:              entity.Date(task_data.DueDate),
		SchtirlichHumorescue: entity.Text(task_data.SchtirlichHumorescue),
		Labels:               toEntityLabels(task_data.Labels),
		Status:               entity.Status(task_data.Status),
	}

	err = s.r.Create(task)

	return *task, err
}

func (t *UpdateTaskRequest) Validate() error {
	if !(t.Name.validate() &&
		t.Description.validate() &&
		t.DueDate.validate() &&
		t.SchtirlichHumorescue.validate() &&
		t.Labels.validate() &&
		t.Status.validate()) {
		return errors.ErrValidation
	}

	return nil
}

// Update modifies task using request
func (s service) Update(request *UpdateTaskRequest) (entity.Task, error) {
	err := request.Validate()

	if err != nil {
		return entity.Task{}, fmt.Errorf("%w: %v", errors.ErrValidation, err)
	}

	task, err := s.r.GetById(request.TaskId)

	if err != nil {
		return entity.Task{}, err
	}

	or := func(ss ...string) string {
		for _, s := range ss {
			if s != "" {
				return s
			}
		}
		return ""
	}

	task.Name = entity.Name(or(string(request.Name), string(task.Name)))
	task.Description = entity.Text(or(string(request.Description), string(task.Description)))
	task.DueDate = entity.Date(or(string(request.DueDate), string(task.DueDate)))
	task.SchtirlichHumorescue = entity.Text(or(string(request.SchtirlichHumorescue), string(task.SchtirlichHumorescue)))
	task.Labels = toEntityLabels(request.Labels)
	task.Status = entity.Status(or(string(request.Status), string(task.Status)))

	err = s.r.Update(&task)

	if err != nil {
		return entity.Task{}, err
	}

	task, err = s.r.GetById(task.Id)

	return task, err
}

// Delete removes task with specified id
func (s service) Delete(task_id int64) (entity.Task, error) {
	task, err := s.r.GetById(task_id)
	if err != nil {
		return entity.Task{}, err
	}

	err = s.r.Delete(task_id)

	return task, err
}
