package ping

import (
	"net/http"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers sets up routing of the HTTP handlers.
func RegisterHandlers(mux *httprouter.Router, logger log.Logger) {
	mux.GET("/ping", accesslog.Log(handleGet, logger))
}

func handleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusTeapot)
}
