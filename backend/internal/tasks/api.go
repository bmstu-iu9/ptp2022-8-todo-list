package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers registers handlers for Tasks API methods
func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.GET("/users/:user_id/tasks", accesslog.Log(errors.Handle(res.handleGet, res.logger), res.logger))
	mux.GET("/users/:user_id/tasks/:task_id", accesslog.Log(errors.Handle(res.handleGetById, res.logger), res.logger))
	mux.PUT("/users/:user_id/tasks/:task_id", accesslog.Log(errors.Handle(res.handlePut, res.logger), res.logger))
	mux.PATCH("/users/:user_id/tasks/:task_id", accesslog.Log(errors.Handle(res.handlePatch, res.logger), res.logger))
	mux.DELETE("/users/:user_id/tasks/:task_id", accesslog.Log(errors.Handle(res.handleDelete, res.logger), res.logger))
	mux.POST("/users/:user_id/tasks/:task_id/complete", accesslog.Log(errors.Handle(res.handleComplete, res.logger), res.logger))
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
	userId, err := strconv.Atoi(p.ByName("user_id"))
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	taskId, err := strconv.Atoi(p.ByName("task_id"))
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	task, err := res.service.GetById(int64(userId), int64(taskId))

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

func (res *resource) handlePut(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	userId, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	taskId, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	request := SetTaskRequest{UserId: int64(userId), TaskId: int64(taskId)}
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyDecode, err)
	}

	task, err := res.service.Set(&request)

	if err != nil {
		return err
	}

	switch request.Mode {
	case CREATE:
		w.WriteHeader(http.StatusCreated)
	case REWRITE:
		err = json.NewEncoder(w).Encode(&task)

		if err != nil {
			return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
		}

		w.WriteHeader(http.StatusNoContent)
	}

	return nil
}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	userId, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	taskId, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	request := UpdateTaskRequest{UserId: int64(userId), TaskId: int64(taskId)}
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyDecode, err)
	}

	task, err := res.service.Update(&request)

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
	userId, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	taskId, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	_, err = res.service.Delete(int64(userId), int64(taskId))

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (res *resource) handleComplete(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	userId, err := strconv.Atoi(p.ByName("user_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	taskId, err := strconv.Atoi(p.ByName("task_id"))

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	task, err := res.service.GetById(int64(userId), int64(taskId))

	if err != nil {
		return err
	}

	request := UpdateTaskRequest{
		UserId: task.UserId,
		TaskId: task.Id,
		Status: entity.COMPLETED,
	}

	task, err = res.service.Update(&request)

	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(&task)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
	}

	return nil
}
