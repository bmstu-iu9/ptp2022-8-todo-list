package users

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// Create saves new user in storage.
	Create(user *entity.User) error
	// Get returns User with specified id.
	Get(id int) (entity.User, error)
	// Delete removes User with specified id.
	Delete(id int) error
	// Update modifies User.
	Update(user *entity.User) error
}

type repository struct {
	db *sql.DB
	logger *log.Logger
}

func NewRepository(db *sql.DB, logger *log.Logger) repository {
	return repository{db, logger}
}

func (repo *repository) Create(user *entity.User) error {
	return nil
}

func (repo *repository) Get(id int) (entity.User, error) {
	return entity.User{}, nil
}

func (repo *repository) Delete(id int) error {
	return nil
}

func (repo *repository) Update(user *entity.User) error {
	return nil
}
