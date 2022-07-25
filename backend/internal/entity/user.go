package entity

// A User represents single API user.
type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
