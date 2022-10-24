package entity

// A User represents single API user.
type User struct {
	Id       int64    `json:"id"`
	Email    Email    `json:"email"`
	Nickname Nickname `json:"nickname"`
	Password Password `json:"password"`
}

type (
	// Email represents email.
	Email string
	// Nickname represents user's nickname.
	Nickname string
	// Password represents password.
	Password string
)

type UserDto struct {
	Id       int64    `json:"id"`
	Email    Email    `json:"email"`
	Nickname Nickname `json:"nickname"`
}

func NewUserDto(user User) UserDto {
	return UserDto{
		Id:       user.Id,
		Email:    user.Email,
		Nickname: user.Nickname,
	}
}
