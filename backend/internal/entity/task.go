package entity

// Task status enumeration
const (
	IN_PROGRESS = "in progress"
	DONE        = "done"
	OUTDATED    = "outdated"
)

// Task represents task at all layers
type Task struct {
	Id                   int64       `json:"id"`
	UserId               int64       `json:"userId"`
	Name                 string      `json:"name"`
	Description          string      `json:"description,omitempty"`
	CreatedOn            string      `json:"createdOn"`
	DueDate              string      `json:"dueDate"`
	SchtirlichHumorescue string      `json:"schtirlichHumorescue,omitempty"`
	Labels               []TaskLabel `json:"labels"`
	Status               string      `json:"status"`
}

// TaskLabel represents label for task
type TaskLabel struct {
	Id     int64  `json:"id"`
	TaskId int64  `json:"-"`
	Name   string `json:"text"`
	Color  string `json:"color"`
}
