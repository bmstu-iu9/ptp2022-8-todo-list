package items

import (
	"encoding/json"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
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

	stateFilter := strToItemStateArr(strings.Split(r.URL.Query().Get("state"), ","))
	if stateFilter[0] == "" {
		defaultStateFilter := []entity.ItemState{"equipped", "inventoried", "store"}
		itemFilter.StateFilter = defaultStateFilter
	} else {
		itemFilter.StateFilter = stateFilter
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

	mux.GET("/user/:userId/items", accesslog.Log(errors.Handle(res.handleGetAll, logger), logger))
	mux.GET("/user/:userId/items/:itemId", accesslog.Log(errors.Handle(res.handleGetOne, logger), logger))
	mux.PATCH("/user/:userId/items/:itemId", accesslog.Log(errors.Handle(res.handlePatch, logger), logger))
}

type resource struct {
	service Service
	logger  log.Logger
}

func (res *resource) handleGetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	userId, err := getUserId(p)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}
	filters := NewFilter(r)
	items, err := res.service.GetAll(userId, filters)
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
	}
	return nil
}

func (res *resource) handleGetOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	userId, err := getUserId(p)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}
	itemId, err := getItemId(p)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}
	item, err := res.service.GetOne(userId, itemId)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrNotFound, err)
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
	}
	return nil
}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	userId, err := getUserId(p)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}
	itemId, err := getItemId(p)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}
	data := UpdateItemStateRequest{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyDecode, err)
	}
	item, err := res.service.UpdateItemState(userId, itemId, &data)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrNotFound, err)
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
	}
	return nil
}

func getUserId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("userId"))
}

func getItemId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("itemId"))
}
