package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers(mux *httprouter.Router, service Service) {
	res := resource{service}

	mux.GET("/users/:user_id/tasks", res.handleGet)
	mux.GET("/users/:user_id/tasks/:task_id", res.handleGetById)
	mux.POST("/users/:user_id/tasks/:task_id", res.handlePost)
	mux.PATCH("/users/:user_id/tasks/:task_id", res.handlePatch)
	mux.DELETE("/users/:user_id/tasks/:task_id", res.handleDelete)
}

type resource struct {
	service Service
}

func (res *resource) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tasks, err := res.service.Get(int64(id))
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (res *resource) handleGetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user_id, err := strconv.Atoi(p.ByName("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task_id, err := strconv.Atoi(p.ByName("task_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err := res.service.GetById(int64(user_id), int64(task_id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (res *resource) handlePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {w.WriteHeader(501)}
func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {w.WriteHeader(501)}
func (res *resource) handleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {w.WriteHeader(501)}