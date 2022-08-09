package errors

import (
	"errors"
	"net/http"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
)

// Handler type represents handler that can return error. It's intended to be
// wrapped in error.Handle handler
type Handler func(http.ResponseWriter, *http.Request, httprouter.Params) error

// Handle creates middleware for handling errors and panics encountered during request handling
func Handle(handler Handler, logger log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				logger.Info(err)
			}
		}()

		err := handler(w, r, p)
		logger.Info(err)
		if err != nil {
			errorResponse(w, err, logger)
		}
	}
}

func errorResponse(w http.ResponseWriter, err error, logger log.Logger) {
	switch {
	case errors.Is(err, ErrNotFound):
		NotFound(w, logger)
	case errors.Is(err, ErrValidation) ||
		errors.Is(err, ErrBodyDecode) || errors.Is(err, ErrLoginFailed) || errors.Is(err, ErrLogoutFailed):
		BadRequest(w, logger)
	case errors.Is(err, ErrWrongPassword):
		Forbidden(w, logger)
	case errors.Is(err, ErrUnauthorized):
		Unauthorized(w, logger)
	default:
		UnexpectedError(w, logger)
	}
}
