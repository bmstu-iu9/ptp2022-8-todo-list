package db

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var fs embed.FS

var (
	host     = config.Get("DB_HOST")
	port     = config.Get("DB_PORT")
	user     = config.Get("DB_USER")
	dbName   = config.Get("DB_NAME")
	password = config.Get("DB_PASSWORD")
	sslMode  = config.Get("DB_SSL_MODE")
)

// New returns new database connection.
func New(logger log.Logger) (*sql.DB, error) {
	params := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbName, password, sslMode)
	logger.Debug("Connecting to DB:", params, "...")
	// TODO: сделать попытки переподключения
	db, err := sql.Open("postgres", params)
	if err != nil {
		return nil, err
	}
	logger.Debug("...done")
	err = runMigrations(db, logger)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewForTest returns new database connection for tests.
func NewForTest(logger log.Logger) (*sql.DB, error) {
	db, err := New(logger)
	if err != nil {
		return db, err
	}
	err = populateTestData(db, logger)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func runMigrations(db *sql.DB, logger log.Logger) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return err
	}
	logger.Debug("Executing migrations...")
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	logger.Debug("...done")
	return nil
}

func populateTestData(db *sql.DB, logger log.Logger) error {
	_, err := db.Exec(`
TRUNCATE users, items RESTART IDENTITY CASCADE;

INSERT INTO users (email, nickname, password) VALUES
('test@example.com', 'test', 'Test123Test'),
('test2@example.com', 'test2', 'Test123Test');

INSERT INTO items (name, image_src, image_for_hero, description, price,
item_category, item_rarity, armor, damage) VALUES
('testItem1', 'test.png', 'test.png', 'test1', 65, 'armor', 'rare', 10, 10),
('testItem2', 'test2.png', 'test2.png', 'test2', 62, 'weapon', 'epic', 0, 5),
('testItem3', 'test3.png', 'test3.png', 'test3', 69, 'weapon', 'legendary', 5, 0);

INSERT INTO inventory (user_id, item_id, item_state) VALUES
(1, 1, 'inventoried'),
(2, 1, 'equipped'),
(1, 2, 'equipped');
`)
	return err
}
