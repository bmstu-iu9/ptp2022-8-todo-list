package errors

import (
	"encoding/json"
	"net/http"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

// Problem represents error description. Conforms RFC7807
type Problem struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Status   int    `json:"status,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func NotFound(w http.ResponseWriter, logger log.Logger) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Problem{
		Title:  "Not found",
		Status: http.StatusNotFound,
	})
	if err != nil {
		logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UnexpectedError(w http.ResponseWriter, logger log.Logger) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Problem{
		Title:  "Unexpected error",
		Status: http.StatusInternalServerError,
	})
	if err != nil {
		logger.Info(err)
	}
}

func BadRequest(w http.ResponseWriter, logger log.Logger) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Problem{
		Title:  "Bad request",
		Status: http.StatusBadRequest,
	})
	if err != nil {
		logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Forbidden(w http.ResponseWriter, logger log.Logger) {
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Problem{
		Title:  "Forbidden",
		Status: http.StatusForbidden,
	})
	if err != nil {
		logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
