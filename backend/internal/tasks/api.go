package tasks

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

// RegisterHandlers registers handlers for Tasks API methods
func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.GET("/users/:user_id/tasks", accesslog.Log(errors.Handle(res.handleGet, res.logger), res.logger))
	mux.GET("/users/:user_id/tasks/:task_id", accesslog.Log(errors.Handle(res.handleGetById, res.logger), res.logger))
	mux.POST("/users/:user_id/tasks", accesslog.Log(errors.Handle(res.handlePost, res.logger), res.logger))
	mux.PATCH("/users/:user_id/tasks/:task_id", accesslog.Log(errors.Handle(res.handlePatch, res.logger), res.logger))
	mux.DELETE("/users/:user_id/tasks/:task_id", accesslog.Log(errors.Handle(res.handleDelete, res.logger), res.logger))
}

type resource struct {
	service Service
	logger  log.Logger
}

func (res *resource) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	id, err := strconv.Atoi(p.ByName("user_id"))
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	// TODO: auth check
	if id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	tasks, err := res.service.Get(int64(id))

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
	}

	return nil
}

func (res *resource) handleGetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	user_id, err := strconv.Atoi(p.ByName("user_id"))
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	// TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	task_id, err := strconv.Atoi(p.ByName("task_id"))
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	task, err := res.service.GetById(int64(task_id))

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
	}

	return nil
}

func (res *resource) handlePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	user_id, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	// TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	task_req := CreateTaskRequest{UserId: int64(user_id)}
	err = json.NewDecoder(r.Body).Decode(&task_req)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyDecode, err)
	}

	task, err := res.service.Create(&task_req)

	if err != nil {
		return err
	}

	w.Header().Set("Location", fmt.Sprintf("%v/users/%v/tasks/%v", config.Get("API_SERVER"), strconv.Itoa(user_id), strconv.Itoa(int(task.Id))))
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	user_id, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	//TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user_id")
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	task_id, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	task_req := UpdateTaskRequest{TaskId: int64(task_id)}
	err = json.NewDecoder(r.Body).Decode(&task_req)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyDecode, err)
	}

	task, err := res.service.Update(&task_req)

	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(&task)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
	}

	return nil
}

func (res *resource) handleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	user_id, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	// TODO: auth check
	if user_id == 0 {
		res.logger.Debug("Zero user id")
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	task_id, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	_, err = res.service.Delete(int64(task_id))

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
