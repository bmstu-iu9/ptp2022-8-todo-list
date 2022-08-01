package items

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

// Repository encapsulates the logic to access items from the data source.
type Repository interface {
	// GetAll returns all items in the application.
	GetAll() ([]entity.Item, error)
	// GetOne returns user's item with specified id.
	GetOne(userId, itemId int) (entity.Item, error)
	// Update modifies the user's item status with specified id.
	Update(userId int, item *entity.Item) error
}

// repository persists items in database.
type repository struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepository creates a new item's repository.
func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// GetAll reads all items from database.
func (repo repository) GetAll() ([]entity.Item, error) {
	rows, err := repo.db.Query("SELECT * FROM items ORDER BY id")
	if err != nil {
		return nil, err
	}
	items := make([]entity.Item, 0)
	for rows.Next() {
		curItem := entity.Item{}
		err = rows.Scan(&curItem.ItemId, &curItem.ItemName, &curItem.ImageSrc, &curItem.Description,
			&curItem.Price, &curItem.Category, &curItem.Rarity)
		if err != nil {
			return nil, err
		}
		items = append(items, curItem)
	}
	return items, nil
}

func (repo repository) isItemInInventory(userId, itemId int) (status bool, err error) {
	row, err := repo.db.Query("SELECT EXISTS "+
		"(SELECT * FROM inventory WHERE user_id = $1 AND item_id = $2)",
		userId, itemId)
	if err != nil {
		return false, err
	}
	row.Next()
	err = row.Scan(&status)
	return status, err
}

// GetOne reads the item with specified id owned by the user with the specified id from database.
func (repo repository) GetOne(userId, itemId int) (entity.Item, error) {
	status, err := repo.isItemInInventory(userId, itemId)
	if err != nil {
		return entity.Item{}, err
	}
	if !status {
		return entity.Item{}, errors.New(fmt.Sprintf("The item with id = %d does not belong "+
			"to user with id =%d", itemId, userId))
	}

	row, err := repo.db.Query("SELECT items.id, name, image_src,"+
		" description, price, item_category, item_rarity, inventory.item_state FROM items INNER JOIN inventory ON items.id = $1"+
		"AND inventory.item_id = $2 AND inventory.user_id = $3", itemId, itemId, userId)
	if err != nil {
		return entity.Item{}, err
	}
	row.Next()
	curItem := entity.Item{}
	err = row.Scan(&curItem.ItemId, &curItem.ItemName, &curItem.ImageSrc, &curItem.Description,
		&curItem.Price, &curItem.Category, &curItem.Rarity, &curItem.ItemState)
	return curItem, err
}

// Update changes the item's state in database.
func (repo repository) Update(userId int, item *entity.Item) error {
	status, err := repo.isItemInInventory(userId, item.ItemId)
	if err != nil {
		return err
	}
	if !status {
		return errors.New(fmt.Sprintf("The item with id = %d does not belong "+
			"to user with id =%d", item.ItemId, userId))
	}

	_, err = repo.db.Exec("UPDATE inventory SET item_state = $1 WHERE item_id = $2 AND user_id =$3",
		item.ItemState, item.ItemId, userId)
	return err
}
