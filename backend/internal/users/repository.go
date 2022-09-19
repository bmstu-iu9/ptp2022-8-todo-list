package users

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// Create saves new user in storage.
	Create(user *entity.User) error
	// Get returns User with specified id.
	Get(id int64) (entity.User, error)
	// Delete removes User with specified id.
	Delete(id int64) error
	// Update modifies User.
	Update(user *entity.User) error
	// InitUserInventory init the list of items for user with spec id.
	InitUserInventory(id int64) error
	// CleanUserInventory deletes the list of items for user with spec id.
	CleanUserInventory(id int64) error
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

// repository persists users in database.
type repository struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepository creates a new users repository.
func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Create saves a new user in repository and sets an id field of user argument
// to the id of saved user.
func (repo repository) Create(user *entity.User) error {
	err := repo.db.QueryRow("INSERT INTO users(email, nickname, password)"+
		"VALUES ($1, $2, $3) RETURNING id", user.Email, user.Nickname, user.Password).
		Scan(&user.Id)
	return wrapSql(err)
}

// InitUserInventory init the list of items for user with spec id.
func (repo repository) InitUserInventory(id int64) error {
	_, err := repo.db.Exec("INSERT INTO inventory (user_id, item_id) SELECT users.id, items.id "+
		"FROM users INNER JOIN items ON users.id = $1 ", id)
	return wrapSql(err)
}

// CleanUserInventory deletes the list of items for user with spec id.
func (repo repository) CleanUserInventory(id int64) error {
	_, err := repo.db.Exec("DELETE FROM inventory WHERE user_id = $1 ", id)
	return wrapSql(err)
}

// Get reads the user with specified id from database.
func (repo repository) Get(id int64) (entity.User, error) {
	user := entity.User{}
	err := repo.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.Nickname, &user.Password)
	return user, wrapSql(err)
}

// Delete removes a user with specified id from database.
func (repo repository) Delete(id int64) error {
	_, err := repo.db.Exec("DELETE FROM users WHERE id = $1", id)
	return wrapSql(err)
}

// Update saves changes to a user from database.
func (repo repository) Update(user *entity.User) error {
	_, err := repo.db.Exec("UPDATE users SET email = $1, nickname = $2,"+
		"password = $3 WHERE id = $4", user.Email, user.Nickname, user.Password, user.Id)
	return wrapSql(err)
}
