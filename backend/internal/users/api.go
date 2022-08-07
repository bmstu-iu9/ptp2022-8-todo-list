package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
)

// RefisterHandlers registers handlers for Users API methods.
func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.POST("/users", accesslog.Log(errors.Handle(res.handlePost, logger), logger))
	mux.GET("/users/:id", accesslog.Log(errors.Handle(res.handleGet, logger), logger))
	mux.DELETE("/users/:id", accesslog.Log(errors.Handle(res.handleDelete, logger), logger))
	mux.PATCH("/users/:id", accesslog.Log(errors.Handle(res.handlePatch, logger), logger))
}

type resource struct {
	service Service
	logger log.Logger
}

func (res *resource) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	id, err := getId(p)
	if err != nil {
		return wrapPath(err)
	}
	user, err := res.service.Get(id)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return wrapEncode(err)
	}
	return nil
}

func (res *resource) handlePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	data := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return wrapDecode(err)
	}
	user, err := res.service.Create(&data)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%v/users/%v", config.Get("API_SERVER"), strconv.FormatInt(int64(user.Id), 10)))
	return nil
}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	id, err := getId(p)
	if err != nil {
		return wrapPath(err)
	}
	data := UpdateUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return wrapDecode(err)
	}
	user, err := res.service.Update(id, &data)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return wrapEncode(err)
	}
	return nil
}

func (res *resource) handleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	id, err := getId(p)
	if err != nil {
		return wrapPath(err)
	}
	_, err = res.service.Delete(id)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func getId(p httprouter.Params) (int64, error) {
	id, err := strconv.Atoi(p.ByName("id"))
	return int64(id), err
}

func wrapDecode(err error) error {
	return fmt.Errorf("%w: %v", errors.ErrBodyDecode, err)
}

func wrapEncode(err error) error {
	return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
}

func wrapPath(err error) error {
	return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
}
