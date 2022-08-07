package auth

import (
	"database/sql"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"log"
)

type Repository interface {
	GetToken(refreshToken string) (DbToken, error)
	UpdateToken(userId int, refreshToken string) error
	CreateToken(userId int, refreshToken string) error
	DeleteToken(refreshToken string) error
	GetUser(email string, userId int) (entity.User, error)
}

type repository struct {
	db     *sql.DB
	logger *log.Logger
}

func NewRepository(db *sql.DB, logger *log.Logger) Repository {
	return repository{db, logger}
}

func (repo repository) GetUser(email string, userId int) (entity.User, error) {
	var row *sql.Rows
	var err error
	if userId == -1 {
		row, err = repo.db.Query("SELECT * FROM users WHERE email = $1", email)
	} else {
		row, err = repo.db.Query("SELECT * FROM users WHERE id = $1", userId)
	}
	if err != nil {
		return entity.User{}, err
	}
	defer row.Close()
	user := entity.User{}
	row.Next()
	err = row.Scan(&user.Id, &user.Email, &user.Nickname, &user.Password)
	return user, err
}

func (repo repository) GetToken(refreshToken string) (DbToken, error) {
	row, err := repo.db.Query("SELECT * FROM tokens WHERE token=$1", refreshToken)
	if err != nil {
		return DbToken{}, err
	}
	var token DbToken
	row.Next()
	err = row.Scan(&token.userId, &token.refreshToken)
	return token, err
}

func (repo repository) UpdateToken(userId int, refreshToken string) error {
	_, err := repo.db.Exec("UPDATE tokens SET token=$1 WHERE user_id=$2", refreshToken, userId)
	return err
}

func (repo repository) DeleteToken(refreshToken string) error {
	_, err := repo.db.Exec("DELETE FROM tokens  WHERE token=$1", refreshToken)
	return err
}

func (repo repository) CreateToken(userId int, refreshToken string) error {
	_, err := repo.db.Exec("INSERT INTO tokens (user_id, token) VALUES ($1,$2)", userId, refreshToken)
	return err
}
