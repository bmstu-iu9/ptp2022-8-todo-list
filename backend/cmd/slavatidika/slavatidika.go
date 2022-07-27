package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/logger"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/ping"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/users"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.New()
	mux := httprouter.New()

	ping.RegisterHandlers(mux)
	users.RegisterHandlers(
		mux,
		users.NewService(users.NewRepository(db, logger)),
		logger)

	server := http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.Host, config.Port),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
