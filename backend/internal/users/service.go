package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"regexp"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

// Service encapsulates usecase logic for users.
type Service interface {
	Get(id int64) (entity.UserDto, error)
	Delete(id int64) (entity.UserDto, error)
	Create(input *CreateUserRequest) (entity.UserDto, error)
	Update(id int64, input *UpdateUserRequest) (entity.UserDto, error)
}

// NewUser represents the data for creating new User.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func validateField(field string, minLen, maxLen int, regex string) bool {
	if matched, _ := regexp.MatchString(regex, field); !matched ||
		len(field) < minLen || len(field) > maxLen {
		return false
	}
	return true
}

// Validate validates the CreateUserRequest fields.
func (req *CreateUserRequest) Validate() error {
	switch {
	case !validateField(req.Email, 1, 200, `^[^\s@]+@[^\s@]+\.[^\s@]+$`):
		return errors.New("Wrong email")
	case !validateField(req.Nickname, 4, 20, `^([a-z\d]+-)*[a-z\d]+$`):
		return errors.New("Wrong nickname")
	case !validateField(req.Password, 8, 100, `^[A-Za-z0-9]\w{8,}$`):
		return errors.New("Wrong password")
	default:
		return nil
	}
}

// UpdateUserRequest represents the data for modifing User.
// Fields Email, Nickname and NewPassword are optional.
type UpdateUserRequest struct {
	Email           *string `json:"email"`
	Nickname        *string `json:"nickname"`
	NewPassword     *string `json:"newPassword"`
	CurrentPassword string  `json:"currentPassword"`
}

// Validate validates the UpdateUserRequest fields.
func (req *UpdateUserRequest) Validate() error {
	switch {
	case req.Email != nil &&
		!validateField(*req.Email, 1, 200, `^[^\s@]+@[^\s@]+\.[^\s@]+$`):
		return errors.New("Wrong email")
	case req.Nickname != nil &&
		!validateField(*req.Nickname, 4, 20, `^([a-z\d]+-)*[a-z\d]+$`):
		return errors.New("Wrong nickname")
	case req.NewPassword != nil &&
		!validateField(*req.NewPassword, 8, 100, `^[A-Za-z0-9]\w{8,}$`):
		return errors.New("Wrong password")
	default:
		return nil
	}
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
func (s service) Create(input *CreateUserRequest) (entity.UserDto, error) {
	err := input.Validate()
	if err != nil {
		return entity.UserDto{}, err
	}

	entityUser := entity.User{
		Email:    input.Email,
		Nickname: input.Nickname,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(input.Password))),
	}
	err = s.repo.Create(&entityUser)
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(entityUser), nil
}

// Update modifies User with given id.
func (s service) Update(id int64, input *UpdateUserRequest) (entity.UserDto, error) {
	err := input.Validate()
	if err != nil {
		return entity.UserDto{}, err
	}

	entityUser, err := s.repo.Get(id)
	if err != nil {
		return entity.UserDto{}, err
	}

	if entityUser.Password != input.CurrentPassword {
		return entity.UserDto{}, errors.New("Wrong password")
	}

	if input.Email != nil {
		entityUser.Email = *input.Email
	}
	if input.Nickname != nil {
		entityUser.Nickname = *input.Nickname
	}
	if input.NewPassword != nil {
		entityUser.Password = *input.NewPassword
	}
	err = s.repo.Update(&entityUser)
	if err != nil {
		return entity.UserDto{}, nil
	}
	return entity.NewUserDto(entityUser), nil
}
