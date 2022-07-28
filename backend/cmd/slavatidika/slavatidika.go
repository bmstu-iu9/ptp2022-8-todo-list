package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/ping"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/users"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := log.New()
	db, err := db.New()
	if err != nil {
		logger.Info(err)
		os.Exit(1)
	}
	mux := httprouter.New()

	ping.RegisterHandlers(mux)
	users.RegisterHandlers(
		mux,
		users.NewService(users.NewRepository(db, logger)),
		logger)

	server := http.Server{
		Addr:    fmt.Sprintf("%v:%v",
			config.Get("HTTP_HOST"), config.Get("HTTP_PORT")),
		Handler: mux,
	}

	logger.Info(server.ListenAndServe())
	os.Exit(1)
}
