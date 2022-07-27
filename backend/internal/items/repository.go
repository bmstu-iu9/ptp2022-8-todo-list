package items

import (
	"database/sql"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"log"
)

// Repository encapsulates the logic to access items from the data source.
type Repository interface {
	// GetAll returns all items in the application.
	GetAll() ([]Item, error)
	// GetOne returns user's item with specified id.
	GetOne(user *entity.User, item *entity.Item) (entity.Item, error)
	// Modify modifies the users`s item with specified id.
	Modify(user *entity.User, item *entity.Item) error
}

type repository struct {
	db     *sql.DB
	logger *log.Logger
}

func (repo *repository) GetAll() ([]Item, error) {
	return nil, nil
}

func (repo *repository) GetOne(user *entity.User, item *entity.Item) (entity.Item, error) {
	return entity.Item{}, nil
}

func (repo *repository) Modify(user *entity.User, item *entity.Item) error {
	return nil
}
