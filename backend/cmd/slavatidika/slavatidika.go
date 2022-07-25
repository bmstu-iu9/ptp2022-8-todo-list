package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/ping"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	ping.RegisterHandlers(mux)

	server := http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.Host, config.Port),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
