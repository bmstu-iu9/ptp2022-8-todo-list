package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
)

// RefisterHandlers registers handlers for Users API methods.
func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.POST("/users", accesslog.Log(res.handlePost, logger))
	mux.GET("/users/:id", accesslog.Log(res.handleGet, logger))
	mux.DELETE("/users/:id", accesslog.Log(res.handleDelete, logger))
	mux.PATCH("/users/:id", accesslog.Log(res.handlePatch, logger))
}

type resource struct {
	service Service
	logger log.Logger
}

func (res *resource) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := getId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user, err := res.service.Get(id)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (res *resource) handlePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := res.service.Create(&data)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%v/users/%v", config.Get("API_SERVER"), strconv.FormatInt(int64(user.Id), 10)))
	w.WriteHeader(http.StatusCreated)
}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := getId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data := UpdateUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := res.service.Update(id, &data)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (res *resource) handleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := getId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = res.service.Delete(id)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getId(p httprouter.Params) (int64, error) {
	id, err := strconv.Atoi(p.ByName("id"))
	return int64(id), err
}
