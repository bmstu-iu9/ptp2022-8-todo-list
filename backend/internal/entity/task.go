package entity

// Task status enumeration
const (
	ACTIVE    = "active"
	COMPLETED = "completed"
	ARCHIVED  = "archived"
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

type (
	Name   string
	Text   string
	Date   string
	Labels string
	Status string
	Color  string
)
