package db

import (
	"database/sql"
	"fmt"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	_ "github.com/lib/pq"
)

var (
	host = config.Get("DB_HOST")
	port = config.Get("DB_PORT")
	user = config.Get("DB_USER")
	dbName = config.Get("DB_NAME")
	password = config.Get("DB_PASSWORD")
	sslMode = config.Get("DB_SSL_MODE")
)

func New(logger log.Logger) (*sql.DB, error) {
	params := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			host, port, user, dbName, password, sslMode)
	logger.Debug("Connecting to DB:", params)
	db, err := sql.Open("postgres", params)
	if err != nil {
		return nil, err
	}

	logger.Debug("Creating new table and test user")
	_, err = db.Exec(`
DROP TABLE IF EXISTS users;
CREATE TABLE users (
       id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
       email varchar(255) UNIQUE NOT NULL,
       nickname varchar(45) NOT NULL,
       password varchar(100) NOT NULL
);
INSERT INTO users(email, nickname, password)
VALUES('test@example.com', 'test', 'Test123Test');
`)

	return db, err
}
