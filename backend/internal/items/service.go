package items

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

// UpdateItemStateRequest represents the data for modifing ItemState.
type UpdateItemStateRequest struct {
	ItemState entity.State `json:"item_state"`
}

// Service encapsulates usecase logic for items.
type Service interface {
	GetAll(userId int, filters entity.Filter) ([]entity.Item, error)
	GetOne(userId, itemId int) (entity.Item, error)
	Modify(userId, itemId int, input *UpdateItemStateRequest) (entity.Item, error)
}

type service struct {
	repo Repository
}

// NewService creates a new item service.
func NewService(repo Repository) Service {
	return service{repo}
}

//GetAll returns all items.
func (s service) GetAll(userId int, filters entity.Filter) ([]entity.Item, error) {
	return s.repo.GetAll(userId, filters)
}

// GetOne returns item with specified id owned by user with specified id.
func (s service) GetOne(userId, itemId int) (entity.Item, error) {
	return s.repo.GetOne(userId, itemId)
}

// Modify returns item with specified id with new ItemState.
func (s service) Modify(userId, itemId int, input *UpdateItemStateRequest) (entity.Item, error) {
	entityItem, err := s.repo.GetOne(userId, itemId)
	if err != nil {
		return entity.Item{}, err
	}
	if input.ItemState == entity.Equipped || input.ItemState == entity.Inventoried {
		entityItem.ItemState = input.ItemState
	}
	err = s.repo.Update(userId, &entityItem)
	return entityItem, err
}
