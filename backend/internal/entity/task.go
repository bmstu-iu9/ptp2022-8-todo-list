package entity

const (
	IN_PROGRESS = "in progress"
	DONE 		= "done"
	OUTDATED 	= "outdated"
)

type Task struct {
	Id						int64 		`json:"id"`
	UserId					int64		`json:"userId"`
	Name					string		`json:"name"`
	Description				string		`json:"description,omitempty"`
	CreatedOn				string		`json:"createdOn"`
	DueDate					string		`json:"dueDate"`
	SchtirlichHumorescue 	string		`json:"schtirlichHumorescue,omitempty"`
	Labels					[]TaskLabel `json:"labels"`
	Status					string		`json:"status"`
}  // TODO: complete task struct definition

type TaskLabel struct {
	Id		int64	`json:"-"`
	TaskId 	int64	`json:"-"`
	Name 	string 	`json:"text"`
	Color 	string	`json:"color"`
}