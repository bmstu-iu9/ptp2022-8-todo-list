package items

import (
	"encoding/json"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.GET("/users/:userId/items", accesslog.Log(res.handleGetAll, logger))
	mux.GET("/users/:userId/items/:itemId", accesslog.Log(res.handleGetOne, logger))
	mux.PATCH("/users/:userId/items/:itemId", accesslog.Log(res.handlePatch, logger))
}

type resource struct {
	service Service
	logger  log.Logger
}

func newFilter(r *http.Request) entity.Filter {
	return entity.Filter{
		StateFilter:    entity.State(r.URL.Query().Get("statefilter")),
		RarityFilter:   r.URL.Query().Get("rarityfilter"),
		CategoryFilter: r.URL.Query().Get("categoryfilter"),
	}
}

func (res *resource) handleGetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := getUserId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	filters := newFilter(r)
	items, err := res.service.GetAll(userId, filters)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (res *resource) handleGetOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := getUserId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	itemId, err := getItemId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item, err := res.service.GetOne(userId, itemId)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := getUserId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	itemId, err := getItemId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data := UpdateItemStateRequest{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := res.service.Modify(userId, itemId, &data)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getUserId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("userId"))
}

func getItemId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("itemId"))
}
