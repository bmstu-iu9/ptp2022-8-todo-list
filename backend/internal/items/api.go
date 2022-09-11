package items

import (
	"encoding/json"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func strToItemStateArr(states []string) []entity.ItemState {
	itemState := make([]entity.ItemState, 0)
	for _, state := range states {
		itemState = append(itemState, entity.ItemState(state))
	}
	return itemState
}

// NewFilter creates a new item filter.
func NewFilter(r *http.Request) ItemFilter {
	var itemFilter ItemFilter

	stateFilter := strings.Split(r.URL.Query().Get("state"), ",")
	if stateFilter[0] == "" {
		defaultStateFilter := []entity.ItemState{"equipped", "inventoried", "store"}
		itemFilter.StateFilter = defaultStateFilter
	} else {
		itemFilter.StateFilter = strToItemStateArr(stateFilter)
	}

	rarityFilter := strings.Split(r.URL.Query().Get("rarity"), ",")
	if rarityFilter[0] == "" {
		defaultRarityFilter := []string{"common", "rare", "epic", "legendary"}
		itemFilter.RarityFilter = defaultRarityFilter
	} else {
		itemFilter.RarityFilter = rarityFilter
	}

	categoryFilter := strings.Split(r.URL.Query().Get("category"), ",")
	if categoryFilter[0] == "" {
		defaultCategoryFilter := []string{"armor", "weapon", "pet", "skin"}
		itemFilter.CategoryFilter = defaultCategoryFilter
	} else {
		itemFilter.CategoryFilter = categoryFilter
	}
	return itemFilter
}

// RegisterHandlers registers handlers for Items API methods.
func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.GET("/user/:userId/items", accesslog.Log(res.handleGetAll, logger))
	mux.GET("/user/:userId/items/:itemId", accesslog.Log(res.handleGetOne, logger))
	mux.PATCH("/user/:userId/items/:itemId", accesslog.Log(res.handlePatch, logger))
}

type resource struct {
	service Service
	logger  log.Logger
}

func (res *resource) handleGetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := getUserId(p)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	filters := NewFilter(r)
	items, err := res.service.GetAll(userId, filters)
	if err != nil {
		res.logger.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		res.logger.Info(err)
		return
	}
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
		return
	}
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
	item, err := res.service.UpdateItemState(userId, itemId, &data)
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
}

func getUserId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("userId"))
}

func getItemId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("itemId"))
}
