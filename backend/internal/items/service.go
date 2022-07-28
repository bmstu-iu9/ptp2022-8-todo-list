package items

type Item struct {
	ItemId   int    `json:"ItemId"`
	ItemName string `json:"ItemName"`
}

// Service encapsulates usecase logic for users.
type Service interface {
	GetAll() ([]Item, error)
	GetOne(userId, itemId int) (Item, error)
	Modify(input *UpdateItemRequest) (Item, error)
}

type UpdateItemRequest struct {
	ItemName string `json:"ItemName"`
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) GetAll() ([]Item, error) {
	allItems, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return allItems, nil
}

func (s service) GetOne(userId, itemId int) (Item, error) {
	item, err := s.repo.GetOne(userId, itemId)
	if err != nil {
		return Item{}, err
	}
	return Item(item), nil
}

func (s service) Modify(input *UpdateItemRequest) (Item, error) {
	return Item{}, nil
}
