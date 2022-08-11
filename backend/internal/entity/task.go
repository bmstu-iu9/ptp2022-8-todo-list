package entity

type Task struct {
	Id			int64 	`json:"id"`
	UserId		int64	`json:"userId"`
	Name		string	`json:"name"`
	Description	string	`json:"description,omitempty"`
}  // TODO: complete task struct definition