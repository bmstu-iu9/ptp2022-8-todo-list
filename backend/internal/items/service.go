package items

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
)

// UpdateItemStateRequest represents the data for modifing ItemState.
type UpdateItemStateRequest struct {
	ItemState entity.ItemState `json:"itemState"`
}

// Service encapsulates usecase logic for items.
type Service interface {
	// GetAll returns all items.
	GetAll(userId int, filters ItemFilter) ([]entity.Item, error)
	// GetOne returns item with specified id owned by user with specified id.
	GetOne(userId, itemId int) (entity.Item, error)
	// UpdateItemState returns item with specified id with new ItemState.
	UpdateItemState(userId, itemId int, input *UpdateItemStateRequest) (entity.Item, error)
}

type service struct {
	repo Repository
}

// NewService creates a new item service.
func NewService(repo Repository) Service {
	return service{repo}
}

// GetAll returns all items.
func (s service) GetAll(userId int, filters ItemFilter) ([]entity.Item, error) {
	return s.repo.GetAll(userId, filters)
}

// GetOne returns item with specified id owned by user with specified id.
func (s service) GetOne(userId, itemId int) (entity.Item, error) {
	return s.repo.GetOne(userId, itemId)
}

// UpdateItemState returns item with specified id with new ItemState.
func (s service) UpdateItemState(userId, itemId int, input *UpdateItemStateRequest) (entity.Item, error) {
	entityItem, err := s.GetOne(userId, itemId)
	if err != nil {
		return entity.Item{}, err
	}
	entityItem.State = input.ItemState
	err = s.repo.Update(userId, &entityItem)
	return entityItem, err
}
