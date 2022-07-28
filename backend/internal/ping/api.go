package ping

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers sets up routing of the HTTP handlers.
func RegisterHandlers(mux *httprouter.Router) {
	mux.GET("/ping", ping)
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(418)
}
