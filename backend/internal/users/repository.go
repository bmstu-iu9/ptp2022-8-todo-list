package users

import (
	"database/sql"
	"fmt"
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
	db     *sql.DB
	logger *log.Logger
}

func NewRepository(db *sql.DB, logger *log.Logger) repository {
	return repository{db, logger}
}

func (repo *repository) Create(user *entity.User) error {
	return nil
}

func (repo *repository) Get(id int) (entity.User, error) {
	row, err := repo.db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %d", id))
	if err != nil {
		return entity.User{}, err
	}
	defer row.Close()
	user := entity.User{}
	err = row.Scan(&user.Id, &user.Email, &user.Nickname, user.Password)
	if err != nil {
		return entity.User{}, err
	} else {
		return user, nil
	}
}

func (repo *repository) Delete(id int) error {
	_, err := repo.db.Exec(fmt.Sprintf("DELETE FROM users WHERE id = %d", id))
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Update(user *entity.User) error {
	return nil
}
