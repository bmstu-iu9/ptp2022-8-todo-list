package items

import (
	"database/sql"
	"log"
)

// Repository encapsulates the logic to access items from the data source.
type Repository interface {
	// GetAll returns all items in the application.
	GetAll() ([]Item, error)
	// GetOne returns user's item with specified id.
	GetOne(userId, itemId int) (Item, error)
	// Update modifies the users`s item with specified id.
	Update(item Item) error
}

type repository struct {
	db     *sql.DB
	logger *log.Logger
}

func (repo *repository) GetAll() ([]Item, error) {
	return nil, nil
}

func (repo *repository) GetOne(userId, itemId int) (Item, error) {
	return Item{}, nil
}

func (repo *repository) Update(item Item) error {
	return nil
}
