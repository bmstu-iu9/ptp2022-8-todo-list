package accesslog

import (
	"net/http"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
)

// Log middleware logs handled request.
func Log(handler httprouter.Handle, logger log.Logger) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger.Info(r.Method, r.URL, r.Proto)
		logger.Debug("\nHeader:", r.Header, "\nBody:", r.Body, "\nHost:", r.Host)
		handler(w, r, p)
	}
}
