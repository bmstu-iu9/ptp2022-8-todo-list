package items

import (
	"database/sql"
	"errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

// Repository encapsulates the logic to access items from the data source.
type Repository interface {
	// GetAll returns all items in the application.
	GetAll() ([]entity.Item, error)
	// GetOne returns user's item with specified id.
	GetOne(userId, itemId int) (entity.Item, error)
	// Update modifies the users`s item status with specified id.
	Update(userId int, item *entity.Item) error
}

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

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

func (repo repository) checkUsersItem(userId, itemId int) (status bool, err error) {
	row, err := repo.db.Query("SELECT EXISTS "+
		"(SELECT * FROM users_items WHERE user_id = $1 AND item_id = $2)",
		userId, itemId)
	if err != nil {
		return false, err
	}
	row.Next()
	err = row.Scan(&status)
	return status, err
}

func (repo repository) GetOne(userId, itemId int) (entity.Item, error) {
	status, err := repo.checkUsersItem(userId, itemId)
	if err != nil {
		return entity.Item{}, err
	}
	if !status {
		return entity.Item{}, errors.New("can`t find data in table")
	}

	row, err := repo.db.Query("SELECT items.id, item_name, image_src,"+
		" description, price, category, rarity, users_items.is_equipped FROM items INNER JOIN users_items ON items.id = $1"+
		"AND users_items.item_id = $2 AND users_items.user_id = $3", itemId, itemId, userId)
	if err != nil {
		return entity.Item{}, err
	}
	row.Next()
	curItem := entity.Item{}
	err = row.Scan(&curItem.ItemId, &curItem.ItemName, &curItem.ImageSrc, &curItem.Description,
		&curItem.Price, &curItem.Category, &curItem.Rarity, &curItem.IsEquipped)
	curItem.IsInInventory = true
	return curItem, err
}

func (repo repository) Update(userId int, item *entity.Item) error {
	status, err := repo.checkUsersItem(userId, item.ItemId)
	item.IsInInventory = status
	if err != nil {
		return err
	}

	_, err = repo.db.Exec("UPDATE users_items SET is_equipped = $1 WHERE item_id = $2 AND user_id =$3",
		item.IsEquipped, item.ItemId, userId)
	return err
}
