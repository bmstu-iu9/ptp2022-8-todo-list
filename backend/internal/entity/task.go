package entity

// Task status enumeration
const (
	IN_PROGRESS = "in progress"
	DONE        = "done"
	OUTDATED    = "outdated"
)

// Task represents task at all layers
type Task struct {
	Id                   int64  `json:"id"`
	UserId               int64  `json:"userId"`
	Name                 Name   `json:"name"`
	Description          Text   `json:"description,omitempty"`
	CreatedOn            Date   `json:"createdOn"`
	DueDate              Date   `json:"dueDate"`
	SchtirlichHumorescue Text   `json:"schtirlichHumorescue,omitempty"`
	Labels               Labels `json:"labels"`
	Status               Status `json:"status"`
}

// TaskLabel represents label for task
type TaskLabel struct {
	Id     int64 `json:"id"`
	TaskId int64 `json:"-"`
	Name   Name  `json:"text"`
	Color  Color `json:"color"`
}

type (
	Name   string
	Text   string
	Date   string
	Labels []TaskLabel
	Status string
	Color  string
)
