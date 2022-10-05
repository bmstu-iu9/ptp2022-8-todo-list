package items

import (
	"database/sql"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

type ItemFilter struct {
	StateFilter    []entity.ItemState
	RarityFilter   []string
	CategoryFilter []string
}

// Repository encapsulates the logic to access items from the data source.
type Repository interface {
	// GetAll returns all items in the application.
	GetAll(userId int, filters ItemFilter) ([]entity.Item, error)
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

func wrapSql(err error) error {
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return fmt.Errorf("%w: %v", errors.ErrNotFound, err)
	default:
		return fmt.Errorf("%w: %v", errors.ErrDb, err)
	}
}

func makeStringForSql(filterParams []string) string {
	s := ""
	for i, param := range filterParams {
		if i == len(filterParams)-1 {
			s += "'" + string(param) + "'"
		} else {
			s += "'" + string(param) + "',"
		}
	}
	return s
}

// GetAll reads all items from database.
func (repo repository) GetAll(userId int, filters ItemFilter) ([]entity.Item, error) {
	rows, err := repo.db.Query(`SELECT items.id, name, image_src, image_for_hero,
          description, price, item_category, item_rarity,armor,damage, inventory.item_state 
          FROM items INNER JOIN inventory ON item_rarity IN 
          (`+makeStringForSql(filters.RarityFilter)+`) AND item_category IN   
          (`+makeStringForSql(filters.CategoryFilter)+`) AND inventory.user_id = $1 AND 
		  items.id = inventory.item_id`, userId)
	if err != nil {
		return nil, wrapSql(err)
	}
	defer rows.Close()
	items := make([]entity.Item, 0)
	curItem := entity.Item{}
	for rows.Next() {
		err = rows.Scan(&curItem.Id, &curItem.Name, &curItem.ImageSrc, &curItem.ImageForHero, &curItem.Description,
			&curItem.Price, &curItem.Category, &curItem.Rarity, &curItem.Armor, &curItem.Damage, &curItem.State)
		if err != nil {
			return nil, wrapSql(err)
		}
		for i := 0; i < len(filters.StateFilter); i++ {
			if curItem.State == filters.StateFilter[i] {
				items = append(items, curItem)
			}
		}
	}
	return items, nil
}

// GetOne reads the item with specified id owned by the user with the specified id from database.
func (repo repository) GetOne(userId, itemId int) (entity.Item, error) {
	curItem := entity.Item{}
	err := repo.db.QueryRow("SELECT items.id, name, image_src, image_for_hero,"+
		" description, price, item_category, item_rarity,armor,damage, inventory.item_state FROM items INNER JOIN "+
		"inventory ON items.id = $1 AND inventory.item_id = $2 AND inventory.user_id = $3",
		itemId, itemId, userId).Scan(&curItem.Id, &curItem.Name, &curItem.ImageSrc, &curItem.ImageForHero, &curItem.Description,
		&curItem.Price, &curItem.Category, &curItem.Rarity, &curItem.Armor, &curItem.Damage, &curItem.State)
	return curItem, wrapSql(err)
}

// Update changes the item's state in database.
func (repo repository) Update(userId int, item *entity.Item) error {
	_, err := repo.db.Exec("UPDATE inventory SET item_state = $1 WHERE item_id = $2 AND user_id =$3",
		item.State, item.Id, userId)
	return wrapSql(err)
}
