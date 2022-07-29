package items

import (
	"database/sql"
	"errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
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
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (repo repository) GetAll() ([]Item, error) {
	rows, err := repo.db.Query("SELECT * FROM items ORDER BY item_id")
	if err != nil {
		return nil, err
	}
	items := make([]Item, 0)
	for rows.Next() {
		curItem := Item{}
		err = rows.Scan(&curItem.ItemId, &curItem.ItemName)
		if err != nil {
			return nil, err
		}
		items = append(items, curItem)
	}
	return items, nil
}

func (repo repository) GetOne(userId, itemId int) (Item, error) {
	row, err := repo.db.Query("SELECT EXISTS "+
		"(SELECT * FROM users_and_items WHERE user_id = $1 AND item_id = $2)",
		userId, itemId)
	if err != nil {
		return Item{}, err
	}
	row.Next()
	var checker bool
	row.Scan(&checker)
	if checker == false {
		return Item{}, errors.New("can`t find data in table")
	}

	row, err = repo.db.Query("SELECT * FROM items WHERE item_id = $1", itemId)
	if err != nil {
		return Item{}, err
	}
	row.Next()
	item := Item{}
	err = row.Scan(&item.ItemId, &item.ItemName)
	return item, err
}

func (repo repository) Update(item Item) error {
	_, err := repo.db.Exec("UPDATE items SET item_name = $1 WHERE item_id = $2",
		item.ItemName, item.ItemId)
	return err
}
