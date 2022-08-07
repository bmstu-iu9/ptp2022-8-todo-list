package main

import (
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/auth"
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
	Db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	Logger := logger.New()
	mux := httprouter.New()

	ping.RegisterHandlers(mux)
	users.RegisterHandlers(
		mux,
		users.NewService(users.NewRepository(Db, Logger)),
		Logger)
	auth.RegisterHandlers(
		mux,
		auth.NewService(auth.NewRepository(Db, Logger)),
		Logger)

	server := http.Server{
		Addr: fmt.Sprintf("%v:%v",
			config.Get("HTTP_HOST"), config.Get("HTTP_PORT")),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
