package items

import (
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/validation"
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
	if !validation.ValidateItemState(filters.StateFilter) {
		return nil, fmt.Errorf("%w: wrong state filter name", errors.ErrValidation)
	}
	if !validation.ValidateItemCategory(filters.CategoryFilter) {
		return nil, fmt.Errorf("%w: wrong category filter name", errors.ErrValidation)
	}
	if !validation.ValidateItemRarity(filters.RarityFilter) {
		return nil, fmt.Errorf("%w: wrong rarity filter name", errors.ErrValidation)
	}
	items, err := s.repo.GetAll(userId, filters)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrNotFound, err)
	}
	return items, nil
}

// GetOne returns item with specified id owned by user with specified id.
func (s service) GetOne(userId, itemId int) (entity.Item, error) {
	return s.repo.GetOne(userId, itemId)
}

// UpdateItemState returns item with specified id with new ItemState.
func (s service) UpdateItemState(userId, itemId int, input *UpdateItemStateRequest) (entity.Item, error) {
	if !validation.ValidateItemState([]entity.ItemState{input.ItemState}) {
		return entity.Item{}, fmt.Errorf("%w: wrong state filter name", errors.ErrValidation)
	}
	entityItem, err := s.GetOne(userId, itemId)
	if err != nil {
		return entity.Item{}, err
	}
	entityItem.State = input.ItemState
	err = s.repo.Update(userId, &entityItem)
	return entityItem, err
}
