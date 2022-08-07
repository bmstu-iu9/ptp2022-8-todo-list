package users

import (
	"fmt"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
)

// Service encapsulates usecase logic for users.
type Service interface {
	Get(id int64) (User, error)
	Delete(id int64) (User, error)
	Create(input *CreateUserRequest) (User, error)
	Update(id int64, input *UpdateUserRequest) (User, error)
}

// User represents the data about an API user.
type User struct {
	Id       int64           `json:"id"`
	Email    entity.Email    `json:"email"`
	Nickname entity.Nickname `json:"nickname"`
}

func newUser(entity *entity.User) User {
	return User{
		Id:       entity.Id,
		Email:    entity.Email,
		Nickname: entity.Nickname,
	}
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
func (s service) Get(id int64) (User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return User{}, err
	}
	return newUser(&user), nil
}

// Delete removes User with specified id.
func (s service) Delete(id int64) (User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return User{}, err
	}
	err = s.repo.Delete(id)
	if err != nil {
		return User{}, err
	}
	return newUser(&user), nil
}

// Create creates User from input data.
func (s service) Create(input *CreateUserRequest) (User, error) {
	err := input.Validate()
	if err != nil {
		return User{}, err
	}

	entityUser := &entity.User{
		Email:    input.Email,
		Nickname: input.Nickname,
		Password: input.Password,
	}
	err = s.repo.Create(entityUser)
	if err != nil {
		return User{}, err
	}
	return newUser(entityUser), nil
}

// Update modifies User with given id.
func (s service) Update(id int64, input *UpdateUserRequest) (User, error) {
	err := input.Validate()
	if err != nil {
		return User{}, err
	}

	entityUser, err := s.repo.Get(id)
	if err != nil {
		return User{}, err
	}

	if entityUser.Password != input.CurrentPassword {
		return User{}, errors.ErrWrongPassword
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
		return User{}, nil
	}
	return newUser(&entityUser), nil
}
