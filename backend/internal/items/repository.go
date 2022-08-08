package items

import (
	"database/sql"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

// Repository encapsulates the logic to access items from the data source.
type Repository interface {
	// GetAll returns all items in the application.
	GetAll(userId int, filters entity.Filter) ([]entity.Item, error)
	// GetOne returns user's item with specified id.
	GetOne(userId, itemId int) (entity.Item, error)
	// Update modifies the user's item status with specified id.
	Update(userId int, item *entity.Item) error
	// IsItemInInventory checks if an item is in inventory.
	IsItemInInventory(userId, itemId int) (entity.State, error)
	// Create creates the item in inventory.
	Create(userId, itemId int, state entity.State) error
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

func createSqlQuery(filters entity.Filter) string {
	var sqlQueryRarity, sqlQueryCategory string
	countOfFilters := 0
	if filters.RarityFilter != "" {
		sqlQueryRarity = fmt.Sprintf("WHERE item_rarity = '%s'", filters.RarityFilter)
		countOfFilters++
	}
	if filters.CategoryFilter != "" {
		if countOfFilters == 0 {
			sqlQueryCategory = fmt.Sprintf(" WHERE item_category = '%s'", filters.CategoryFilter)
		}
		if countOfFilters == 1 {
			sqlQueryCategory = fmt.Sprintf(" AND item_category = '%s'", filters.CategoryFilter)
		}
	}
	return sqlQueryRarity + sqlQueryCategory
}

// GetAll reads all items from database.
func (repo repository) GetAll(userId int, filters entity.Filter) ([]entity.Item, error) {
	sqlQueryFilters := createSqlQuery(filters)
	rows, err := repo.db.Query("SELECT * FROM items " + sqlQueryFilters + "ORDER BY id")
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

		status, _ := repo.IsItemInInventory(userId, curItem.ItemId)
		curItem.ItemState = status
		if filters.StateFilter != entity.Unknown {
			if curItem.ItemState == filters.StateFilter {
				items = append(items, curItem)
			}
		} else {
			items = append(items, curItem)
		}
	}

	return items, nil
}

// IsItemInInventory checks if an item is in inventory.
func (repo repository) IsItemInInventory(userId, itemId int) (entity.State, error) {
	row, err := repo.db.Query("SELECT item_state FROM inventory WHERE user_id = $1 AND item_id = $2",
		userId, itemId)
	if err != nil {
		return entity.Unknown, err
	}
	row.Next()
	var status entity.State
	err = row.Scan(&status)
	if err != nil {
		return entity.Store, fmt.Errorf("The item with id = %d does not belong to user with id =%d",
			itemId, userId)
	}
	return status, nil
}

// GetOne reads the item with specified id owned by the user with the specified id from database.
func (repo repository) GetOne(userId, itemId int) (entity.Item, error) {
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
	_, err := repo.db.Exec("UPDATE inventory SET item_state = $1 WHERE item_id = $2 AND user_id =$3",
		item.ItemState, item.ItemId, userId)
	return err
}

// Create creates the item in inventory.
func (repo repository) Create(userId, itemId int, state entity.State) error {
	_, err := repo.db.Exec("INSERT INTO inventory (user_id,item_id,item_state) VALUES ($1,$2,$3)",
		userId, itemId, state)
	return err
}
