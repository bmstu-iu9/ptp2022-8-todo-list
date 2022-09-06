package errors

import (
	"encoding/json"
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

var errorResponses = []struct {
	err  error
	code int
	body Problem
}{
	{
		err:  ErrNotFound,
		code: http.StatusNotFound,
		body: Problem{
			Title:  "Not found",
			Status: http.StatusNotFound,
		},
	},
	{
		err:  ErrDb,
		code: http.StatusInternalServerError,
		body: Problem{
			Title:  "Internal server error",
			Status: http.StatusInternalServerError,
			Detail: "Database error",
		},
	},
	{
		err:  ErrBodyDecode,
		code: http.StatusBadRequest,
		body: Problem{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "Bad request body",
		},
	},
	{
		err:  ErrBodyEncode,
		code: http.StatusInternalServerError,
		body: Problem{
			Title:  "Internal server error",
			Status: http.StatusInternalServerError,
		},
	},
	{
		err:  ErrPathParameter,
		code: http.StatusBadRequest,
		body: Problem{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "Bad path parameter",
		},
	},
	{
		err:  ErrValidation,
		code: http.StatusBadRequest,
		body: Problem{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "Request body parameters validation failed",
		},
	},
	{
		err:  ErrAuth,
		code: http.StatusForbidden,
		body: Problem{
			Title:  "Forbidden",
			Status: http.StatusForbidden,
			Detail: "Wrong login or password",
		},
	},
	{
		err:  ErrNotAllowed,
		code: http.StatusMethodNotAllowed,
		body: Problem{
			Title:  "Method not allowed",
			Status: http.StatusMethodNotAllowed,
		},
	},
}

func errorResponse(w http.ResponseWriter, err error, logger log.Logger) {
	for _, errorResponse := range errorResponses {
		if errors.Is(err, errorResponse.err) {
			w.WriteHeader(errorResponse.code)
			w.Header().Set("Content-Type", "application/problem+json")
			err := json.NewEncoder(w).Encode(errorResponse.body)
			if err != nil {
				logger.Info(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
}
