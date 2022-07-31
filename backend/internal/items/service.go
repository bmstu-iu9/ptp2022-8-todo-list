package items

import "github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"

type UpdateItemRequest struct {
	IsEquipped    *int  `json:"is_equipped"`
	IsInInventory *bool `json:"is_in_inventory"`
}

// Service encapsulates usecase logic for users.
type Service interface {
	GetAll() ([]entity.Item, error)
	GetOne(userId, itemId int) (entity.Item, error)
	Modify(userId, itemId int, input *UpdateItemRequest) (entity.Item, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) GetAll() ([]entity.Item, error) {
	return s.repo.GetAll()
}

func (s service) GetOne(userId, itemId int) (entity.Item, error) {
	return s.repo.GetOne(userId, itemId)
}

func (s service) Modify(userId, itemId int, input *UpdateItemRequest) (entity.Item, error) {
	entityItem, err := s.repo.GetOne(userId, itemId)
	if err != nil {
		return entity.Item{}, err
	}
	if input.IsEquipped != nil {
		entityItem.IsEquipped = *input.IsEquipped
	}
	if input.IsInInventory != nil {
		entityItem.IsInInventory = *input.IsInInventory
	}
	err = s.repo.Update(userId, &entityItem)
	return entityItem, err
}
