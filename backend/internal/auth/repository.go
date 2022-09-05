package auth

import (
	"database/sql"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

type Repository interface {
	// GetToken reads the token from the database.
	GetToken(refreshToken string, userId int) (DbToken, error)
	// UpdateToken updates user`s refresh token in db.
	UpdateToken(userId int, newRefreshToken string) error
	// CreateToken creates a new refresh token in db.
	CreateToken(userId int, refreshToken string) error
	// DeleteToken deletes a refresh token from db.
	DeleteToken(refreshToken string) error
	// GetUser reads the user from the database.
	GetUser(email entity.Email, userId int) (entity.User, error)
	// DeleteDeadUsers deletes non activated accounts from DB
	DeleteDeadUsers() error
}

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// GetUser reads the user from the database.
// It has 2 formats of work.
// 1) When the user's email is specified, it searches the database by email, userId=-1.
// 2) When userId!=-1 the search is carried out by ID, email at the same time you can not specify.
func (repo repository) GetUser(email entity.Email, userId int) (entity.User, error) {
	var (
		row *sql.Rows
		err error
	)
	if userId == -1 {
		row, err = repo.db.Query("SELECT id,email,nickname,password,is_activated FROM users WHERE email = $1",
			email)
	} else {
		row, err = repo.db.Query("SELECT id,email,nickname,password,is_activated FROM users WHERE id = $1",
			userId)
	}
	if err != nil {
		return entity.User{}, err
	}
	defer row.Close()
	user := entity.User{}
	row.Next()
	err = row.Scan(&user.Id, &user.Email, &user.Nickname, &user.Password, &user.IsActivated)
	return user, err
}

// GetToken The function reads the token from the database.
// It has 2 formats of work.
// 1) When the refreshToken is specified, it searches the database by refreshToken, userId=-1.
// 2) When userId!=-1 the search is carried out by ID, refreshToken at the same time you can not specify.
func (repo repository) GetToken(refreshToken string, userId int) (DbToken, error) {
	var (
		row *sql.Rows
		err error
	)
	if userId == -1 {
		row, err = repo.db.Query("SELECT * FROM tokens WHERE token=$1", refreshToken)
	} else {
		row, err = repo.db.Query("SELECT * FROM tokens WHERE user_id=$1", userId)
	}
	if err != nil {
		return DbToken{}, err
	}
	token := DbToken{}
	row.Next()
	err = row.Scan(&token.userId, &token.refreshToken)
	return token, err
}

// UpdateToken updates user`s refresh token in db.
func (repo repository) UpdateToken(userId int, newRefreshToken string) error {
	_, err := repo.db.Exec("UPDATE tokens SET token=$1 WHERE user_id=$2", newRefreshToken, userId)
	return err
}

// DeleteToken deletes a refresh token from db.
func (repo repository) DeleteToken(refreshToken string) error {
	_, err := repo.db.Exec("DELETE FROM tokens WHERE token=$1", refreshToken)
	return err
}

// CreateToken creates a new refresh token in db.
func (repo repository) CreateToken(userId int, refreshToken string) error {
	_, err := repo.db.Exec("INSERT INTO tokens (user_id, token) VALUES ($1,$2)", userId, refreshToken)
	return err
}

func (repo repository) DeleteDeadUsers() error {
	_, err := repo.db.Exec("DELETE FROM users WHERE is_activated = 'false' AND creation_date > NOW()-INTERVAL '1' DAY")
	return err
}
