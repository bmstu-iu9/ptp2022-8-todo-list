package tasks

type TaskRequest interface {}  // TODO: implement task (according to API) structure
type TasksRequest []TaskRequest

type Service interface {
	Get(user_id int64) (TasksRequest, error)
	GetById(user_id int64, task_id int64) (TaskRequest, error)
	Create(user_id int64, task_id int64, task_data TaskRequest) error
	Update(user_id int64, task_id int64, task_data TaskRequest) error
	Delete(user_id int64, task_id int64) error
}

type service struct {
	r *repository
}

func (s *service) Get(user_id int64) (TasksRequest, error) {
	return nil, nil
}

func (s *service) GetById(user_id int64, task_id int64) (TaskRequest, error) {
	return nil, nil
}

func (s *service) Create(user_id int64, task_id int64, task_data TaskRequest) error {
	return nil
}

func (s *service) Update(user_id int64, task_id int64, task_data TaskRequest) error {
	return nil
}

func (s *service) Delete(user_id int64, task_id int64) error {
	return nil
}