package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/items"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/ping"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/tasks"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/users"
)

func main() {
	logger := log.New()
	db, err := db.New(logger)
	if err != nil {
		logger.Info(err)
		os.Exit(1)
	}
	logger.Debug("DB connection established")
	mux := router.New(logger)

	ping.RegisterHandlers(mux, logger)
	users.RegisterHandlers(
		mux,
		users.NewService(users.NewRepository(db, logger)),
		logger)
	tasks.RegisterHandlers(
		mux,
		tasks.NewService(tasks.NewRepository(db, logger)),
		logger)
	items.RegisterHandlers(
		mux,
		items.NewService(items.NewRepository(db, logger)),
		logger)

	address := fmt.Sprintf("%v:%v",
		config.Get("HTTP_HOST"), config.Get("HTTP_PORT"))
	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	logger.Info("Slavatidika launched on", address)
	logger.Info(server.ListenAndServe())
	os.Exit(1)
}
