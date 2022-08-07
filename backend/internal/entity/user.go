package entity

// A User represents single API user.
type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type UserDto struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

func NewUserDto(entity User) UserDto {
	return UserDto{
		Id:       entity.Id,
		Email:    entity.Email,
		Nickname: entity.Nickname,
	}
}
