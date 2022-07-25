package users

import (
	"errors"
	"regexp"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

// Service encapsulates usecase logic for users.
type Service interface {
	Get(id int) (User, error)
	Delete(id int) (User, error)
	Create(input *CreateUserRequest) (User, error)
	Update(id int, input *UpdateUserRequest) (User, error)
}

// User represents the data about an API user.
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
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
func (s service) Get(id int) (User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return User{}, err
	}
	return newUser(&user), nil
}

// Delete removes User with specified id.
func (s service) Delete(id int) (User, error) {
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
func (s service) Update(id int, input *UpdateUserRequest) (User, error) {
	err := input.Validate()
	if err != nil {
		return User{}, err
	}

	entityUser, err := s.repo.Get(id)
	if err != nil {
		return User{}, err
	}

	if entityUser.Password != input.CurrentPassword {
		return User{}, errors.New("Wrong password")
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
