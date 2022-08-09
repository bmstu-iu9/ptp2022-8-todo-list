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

	logger.Debug("Creating new users table and test user")
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

	logger.Debug("Creating new tasks table and test task")
	_, err = db.Exec(`
DROP TABLE IF EXIST tasks CASCADE;
CREATE TYPE status AS ENUM ('in progress', 'done', 'outdated');
CREATE TABLE tasks (
	id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	user_id int FOREIGN KEY REFERENCES users,
	name varchar(255) NOT NULL,
	description text,
	created_on timestamp NOT NULL,
	due_date timestamp NOT NULL,
	schtirlich_humorescue text NOT NULL,
	cur_status status NOT NULL
);
INSERT INTO tasks(user_id, name, description, created_on, due_date, schtirlich_humorescue, cur_status)
VALUES (
	1,  -- user_id
	'Сделать свою ишью',  -- name
	'СДЕЛАЙ ИШЬЮ.'
	'ВОЗЬМИ И СДЕЛАЙ ИШЬЮ.'
	'Возьми и сделай ишью.'
	'Ты хочешь и дальше сидеть в своем долбаном унылом городе.'
	'Или ты хочешь вырваться в свет.'
	'Сделай ишью.'
	'Твои родители и друзья не верят, что ты можешь сделать ишью и сдать практику.'
	'ПОФИГ.'
	'Возьми и сделай ишью.'
	'Твои друзья шлют проект нафиг.'
	'Возьми его за яйца, и сделай ишью.'
	'Хватит искать долбанные отмазы.'
	'Продолжай вкалывать.'
	'И сделай ишью.'
	'ИШЬЮ'
	,  -- description
	LOCALTIMESTAMP,  -- created_on
	LOCALTIMESTAMP,  -- due_date
	E'Подвыпившие Штирлиц и Мюллер вышли из бара.\n'
	'- Давайте снимем девочек, - предложил Штирлиц.\n'
	'- У вас очень доброе сердце - ответил Мюллер. - Но пусть все-таки повисят до утра.',  -- schtirlich_humorescue
	'in progress'
	)

DROP TABLE IF EXIST task_labels CASCADE;
CREATE TABLE task_labels {
	id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	task_id int FOREIGN KEY REFERENCES tasks,
	name text NOT NULL,
	color int NOT NULL
};
INSERT INTO task_labels(task_id, name, color)
VALUES (1, 'Test label', 255);
`)

	return db, err
}
