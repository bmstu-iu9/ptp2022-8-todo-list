package tasks

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

func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.GET("/users/:user_id/tasks", accesslog.Log(res.handleGet, res.logger))
	mux.GET("/users/:user_id/tasks/:task_id", accesslog.Log(res.handleGetById, res.logger))
	mux.POST("/users/:user_id/tasks", accesslog.Log(res.handlePost, res.logger))
	mux.PATCH("/users/:user_id/tasks/:task_id", accesslog.Log(res.handlePatch, res.logger))
	mux.DELETE("/users/:user_id/tasks/:task_id", accesslog.Log(res.handleDelete, res.logger))
}

type resource struct {
	service Service
	logger log.Logger
}

func (res *resource) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("user_id"))
	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: auth check
	if id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := res.service.Get(int64(id))
	
	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// w.WriteHeader(http.StatusOK)
}

func (res *resource) handleGetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user_id, err := strconv.Atoi(p.ByName("user_id"))
	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task_id, err := strconv.Atoi(p.ByName("task_id"))
	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err := res.service.GetById(int64(task_id))

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
}

func (res *resource) handlePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user_id, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task_req := CreateTaskRequest{UserId: int64(user_id)}
	err = json.NewDecoder(r.Body).Decode(&task_req)
	// res.logger.Debug("Labels received: ", task_req.Labels)

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := res.service.Create(&task_req)

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%v/users/%v/tasks/%v", config.Get("API_SERVER"), strconv.Itoa(user_id), strconv.Itoa(int(task.Id))))
	w.WriteHeader(http.StatusCreated)
}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user_id, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task_id, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task_req := UpdateTaskRequest{TaskId: int64(task_id)}
	err = json.NewDecoder(r.Body).Decode(&task_req)

	logger := log.New()
	logger.Debug("Labels to update: ", task_req.Labels)

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := res.service.Update(&task_req)

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&task)

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
}

func (res *resource) handleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user_id, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task_id, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = res.service.Delete(int64(task_id))

	if err != nil {
		res.logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}