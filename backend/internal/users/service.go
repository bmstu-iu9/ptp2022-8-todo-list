package users

import (
	"crypto/md5"
	"fmt"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
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
	Email    entity.Email    `json:"email"`
	Nickname entity.Nickname `json:"nickname"`
	Password entity.Password `json:"password"`
}

// Validate validates the CreateUserRequest fields.
func (req *CreateUserRequest) Validate() error {
	wrap := func(err error) error {
		return fmt.Errorf("%w: %v", errors.ErrValidation, err)
	}

	if err := req.Email.Validate(); err != nil {
		return wrap(err)
	}
	if err := req.Nickname.Validate(); err != nil {
		return wrap(err)
	}
	if err := req.Password.Validate(); err != nil {
		return wrap(err)
	}
	return nil
}

// UpdateUserRequest represents the data for modifing User.
// Fields Email, Nickname and NewPassword are optional.
type UpdateUserRequest struct {
	Email           *entity.Email    `json:"email"`
	Nickname        *entity.Nickname `json:"nickname"`
	NewPassword     *entity.Password `json:"newPassword"`
	CurrentPassword entity.Password  `json:"currentPassword"`
}

// Validate validates the UpdateUserRequest fields.
func (req *UpdateUserRequest) Validate() error {
	wrap := func(err error) error {
		return fmt.Errorf("%w: %v", errors.ErrValidation, err)
	}

	if err := req.Email.Validate(); err != nil {
		return wrap(err)
	}
	if err := req.Nickname.Validate(); err != nil {
		return wrap(err)
	}
	if err := req.NewPassword.Validate(); err != nil {
		return wrap(err)
	}
	return nil
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
		Password: *entity.NewPassword(fmt.Sprintf("%x", md5.Sum([]byte(input.Password)))),
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
		return entity.UserDto{}, errors.ErrWrongPassword
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
