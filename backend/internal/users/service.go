package users

import (
	"crypto/md5"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"net/smtp"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/google/uuid"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/validation"
)

// Service encapsulates usecase logic for users.
type Service interface {
	Get(id int64) (entity.UserDto, error)
	Delete(id int64) (entity.UserDto, error)
	Create(input *CreateUserRequest) (entity.UserDto, error)
	Update(id int64, input *UpdateUserRequest) (entity.UserDto, error)
	Activate(activationLink string) error
}

// User represents the data about an API user.
type User struct {
	Id       int64    `json:"id"`
	Email    Email    `json:"email"`
	Nickname Nickname `json:"nickname"`
}

type (
	// Email represents email.
	Email    entity.Email
	// Nickname represents user's nickname.
	Nickname entity.Nickname
	// Password represents plaintext password.
	Password entity.Password
)

func newEmail(str string) *Email {
	return (*Email)(&str)
}

func newPassword(str string) *Password {
	return (*Password)(&str)
}

func newNickname(str string) *Nickname {
	return (*Nickname)(&str)
}

func (field *Email) validate() bool {
	if field == nil {
		return true
	}
	return validation.ValidateField(string(*field), 1, 200, `^[^\s@]+@[^\s@]+\.[^\s@]+$`)
}

func (field *Nickname) validate() bool {
	if field == nil {
		return true
	}
	return validation.ValidateField(string(*field), 4, 20, `^([a-z\d]+-)*[a-z\d]+$`)
}

func (field *Password) validate() bool {
	if field == nil {
		return true
	}
	return validation.ValidateField(string(*field), 8, 100, `^[A-Za-z0-9]\w{8,}$`)
}

func newUser(entity *entity.User) User {
	return User{
		Id:       entity.Id,
		Email:    Email(entity.Email),
		Nickname: Nickname(entity.Nickname),
	}
}

// NewUser represents the data for creating new User.
type CreateUserRequest struct {
	Email    Email    `json:"email"`
	Nickname Nickname `json:"nickname"`
	Password Password `json:"password"`
}

// validate validates the CreateUserRequest fields.
func (req *CreateUserRequest) validate() error {
	if req.Email.validate() && req.Nickname.validate() && req.Password.validate() {
		return nil
	}
	return errors.ErrValidation
}

// UpdateUserRequest represents the data for modifing User.
// Fields Email, Nickname and NewPassword are optional.
type UpdateUserRequest struct {
	Email           *Email    `json:"email"`
	Nickname        *Nickname `json:"nickname"`
	NewPassword     *Password `json:"newPassword"`
	CurrentPassword Password  `json:"currentPassword"`
}

// validate validates the UpdateUserRequest fields.
func (req *UpdateUserRequest) validate() error {
	if req.Email.validate() && req.Nickname.validate() && req.NewPassword.validate() {
		return nil
	}
	return errors.ErrValidation
}

type service struct {
	repo Repository
}

// NewService creates a new user service.
func NewService(repo Repository) Service {
	return service{repo}
}

// Get returns User with specified id.
func (s service) Get(id int64) (entity.UserDto, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(user), nil
}

// Delete removes User with specified id.
func (s service) Delete(id int64) (entity.UserDto, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return entity.UserDto{}, err
	}
	err = s.repo.Delete(id)
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(user), nil
}

// Create creates User from input data.
func (s service) Create(input *CreateUserRequest) (User, error) {
	err := input.validate()
	if err != nil {
		return entity.UserDto{}, err
	}
	activationLink, _ := uuid.NewUUID()
	entityUser := entity.User{
		Email:    input.Email,
		Nickname: input.Nickname,
		Password: *entity.NewPassword(fmt.Sprintf("%x", md5.Sum([]byte(input.Password)))),
	}
	err = s.repo.Create(&entityUser, activationLink.String())
	if err != nil {
		return entity.UserDto{}, err
	}
	err = sendActivationMail(entityUser.Email, activationLink.String())
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(entityUser), nil
}

// Update modifies User with given id.
func (s service) Update(id int64, input *UpdateUserRequest) (entity.UserDto, error) {
	err := input.validate()
	if err != nil {
		return entity.UserDto{}, err
	}

	entityUser, err := s.repo.Get(id)
	if err != nil {
		return entity.UserDto{}, err
	}

	if Password(entityUser.Password) != input.CurrentPassword {
		return User{}, errors.ErrAuth
	}

	if input.Email != nil {
		entityUser.Email = entity.Email(*input.Email)
	}
	if input.Nickname != nil {
		entityUser.Nickname = entity.Nickname(*input.Nickname)
	}
	if input.NewPassword != nil {
		entityUser.Password = entity.Password(*input.NewPassword)
	}
	err = s.repo.Update(&entityUser)
	if err != nil {
		return entity.UserDto{}, nil
	}
	return entity.NewUserDto(entityUser), nil
}

func (s service) Activate(activationLink string) error {
	err := s.repo.CheckActivationLink(activationLink)
	if err != nil {
		return err
	}
	err = s.repo.UpdateActivationStatus(activationLink)
	if err != nil {
		return err
	}
	return nil
}

func sendActivationMail(email entity.Email, activationLink string) error {
	link := config.Get("API_SERVER_TEST") + "/activate/" + activationLink
	auth := smtp.PlainAuth("", "slavatidika@gmail.com", "ojlakqwiiuvknvcx",
		"smtp.gmail.com")
	subject := "Subject: Activate your account!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf("<html><body><div><h1>For activation your account click the link below</h1>"+
		"<a href='%s'>%s</a></div></body></html>", link, link)
	to := "To: " + string(email) + "\n"
	fmt.Println(to)
	msg := []byte(subject + to + mime + body)
	err := smtp.SendMail("smtp.gmail.com: 587", auth, "slavatidika@gmail.com",
		[]string{string(email)}, msg)
	if err != nil {
		return err
	}
	return nil
}
