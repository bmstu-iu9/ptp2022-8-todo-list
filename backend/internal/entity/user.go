package entity

import "errors"

// A User represents single API user.
type User struct {
	Id          int64    `json:"id"`
	Email       Email    `json:"email"`
	Nickname    Nickname `json:"nickname"`
	Password    Password `json:"password"`
	IsActivated bool     `json:"is_activated"`
}

type (
	// Email represents email.
	Email string
	// Nickname represents user's nickname.
	Nickname string
	// Password represents password.
	Password string
)

// NewEmail creates new email from string.
func NewEmail(str string) *Email {
	return (*Email)(&str)
}

// NewNickname creates new nickanme from string.
func NewNickname(str string) *Nickname {
	return (*Nickname)(&str)
}

// NewPassword creates new password from string.
func NewPassword(str string) *Password {
	return (*Password)(&str)
}

var (
	ErrEmailValidation    = errors.New("email validation failed")
	ErrNicknameValidation = errors.New("nickname validation failed")
	ErrPasswordValidation = errors.New("password validation failed")
)

// Validate validates email.
// Validation is passed when nil pointer presented.
func (email *Email) Validate() error {
	if email == nil ||
		validateField(string(*email), 1, 200, `^[^\s@]+@[^\s@]+\.[^\s@]+$`) {
		return nil
	}
	return ErrEmailValidation
}

// Validate validates nickname.
// Validation is passed when nil pointer presented.
func (nickname *Nickname) Validate() error {
	if nickname == nil ||
		validateField(string(*nickname), 4, 20, `^([a-z\d]+-)*[a-z\d]+$`) {
		return nil
	}
	return ErrNicknameValidation
}

// Validate validates password.
// Validation is passed when nil pointer presented.
func (password *Password) Validate() error {
	if password == nil ||
		validateField(string(*password), 8, 100, `^[A-Za-z0-9]\w{8,}$`) {
		return nil
	}
	return ErrPasswordValidation
}

type UserDto struct {
	Id       int64    `json:"id"`
	Email    Email    `json:"email"`
	Nickname Nickname `json:"nickname"`
}

func NewUserDto(entity User) UserDto {
	return UserDto{
		Id:       entity.Id,
		Email:    entity.Email,
		Nickname: entity.Nickname,
	}
}
