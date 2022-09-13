package items

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

var (
	mux    *httprouter.Router
	logger log.Logger
)

func init() {
	mux = router.New(logger)
	logger = log.New()
	service := NewService(NewMockRerository())
	RegisterHandlers(mux, service, logger)
}

func TestApi(t *testing.T) {
	toJson := func(data interface{}) string {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(data)
		if err != nil {
			panic(err)
		}
		return buf.String()
	}
	notFound := toJson(errors.Problem{Title: "Not found", Status: http.StatusNotFound})
	badRequest := toJson(errors.Problem{Title: "Bad request", Status: http.StatusBadRequest, Detail: "Request body parameters validation failed"})
	tests := []test.ApiTestCase{
		{
			Name: "get all ok", Method: "GET", Url: "/user/1/items",
			WantBody: toJson([]entity.Item{
				{
					Id:       10,
					Name:     "sword",
					Rarity:   "rare",
					Category: "weapon",
					State:    entity.Equipped,
				},
				{
					Id:       6,
					Name:     "head",
					Rarity:   "legendary",
					Category: "armor",
					State:    entity.Equipped,
				}}),
			WantCode: http.StatusOK,
		},
		{
			Name: "get all fail", Method: "GET", Url: "/user/10/items",
			WantBody: notFound,
			WantCode: http.StatusNotFound,
		},
		{
			Name: "get all rarity filter ok", Method: "GET", Url: "/user/1/items?rarity=rare",
			WantBody: toJson([]entity.Item{{Id: 10, Name: "sword", Rarity: "rare", Category: "weapon",
				State: entity.Equipped}}),
			WantCode: http.StatusOK,
		},
		{
			Name: "get all rarity and category filter ok", Method: "GET", Url: "/user/1/items?rarity=rare&category=pet",
			WantBody: toJson(nil),
			WantCode: http.StatusOK,
		},
		{
			Name: "get all category filter ok", Method: "GET", Url: "/user/1/items?category=weapon",
			WantBody: toJson([]entity.Item{{Id: 10, Name: "sword", Rarity: "rare", Category: "weapon",
				State: entity.Equipped}}),
			WantCode: http.StatusOK,
		},
		{
			Name: "get all rarity filter fail", Method: "GET", Url: "/user/1/items?rarity=cringe",
			WantBody: badRequest,
			WantCode: http.StatusBadRequest,
		},
		{
			Name: "get one ok", Method: "GET", Url: "/user/1/items/10",
			WantBody: toJson(entity.Item{Id: 10, Name: "sword", Rarity: "rare", Category: "weapon",
				State: entity.Equipped}),
			WantCode: http.StatusOK,
		},
		{
			Name: "get one fail user id", Method: "GET", Url: "/user/2/items/10",
			WantBody: notFound,
			WantCode: http.StatusNotFound,
		},
		{
			Name: "get one fail item id", Method: "GET", Url: "/user/1/items/3",
			WantBody: notFound,
			WantCode: http.StatusNotFound,
		},
		{
			Name: "update ok", Method: "PATCH", Url: "/user/1/items/10",
			Body: toJson(UpdateItemStateRequest{ItemState: entity.Inventoried}),
			WantBody: toJson(entity.Item{Id: 10, Name: "sword", Rarity: "rare", Category: "weapon",
				State: entity.Inventoried}),
			WantCode: http.StatusOK,
		},
	}
	test.Endpoint(t, tests, mux)
}

type mockRepository struct {
	data   []entity.Item
	userId int
}

func (m mockRepository) GetAll(userId int, filters ItemFilter) ([]entity.Item, error) {
	if m.userId != userId {
		return nil, fmt.Errorf("no user")
	}
	var ans []entity.Item
	var f bool
	if len(filters.RarityFilter) <= 2 && filters.RarityFilter != nil {
		f = true
		ans = make([]entity.Item, 0)
		for _, item := range m.data {
			for _, rarity := range filters.RarityFilter {
				if item.Rarity == rarity {
					ans = append(ans, item)
				}
			}
		}
	}
	if len(filters.CategoryFilter) <= 2 && filters.CategoryFilter != nil {
		var afterRarityCheck []entity.Item
		f = true
		if ans != nil {
			afterRarityCheck = ans
			ans = nil
		} else {
			afterRarityCheck = m.data
		}
		for _, item := range afterRarityCheck {
			for _, category := range filters.CategoryFilter {
				if item.Category == category {
					ans = append(ans, item)
				}
			}
		}
	}
	if f {
		return ans, nil
	}
	return m.data, nil
}

func (m mockRepository) GetOne(userId, itemId int) (entity.Item, error) {
	if userId != m.userId {
		return entity.Item{}, fmt.Errorf("no user")
	}
	for _, item := range m.data {
		if item.Id == itemId {
			return item, nil
		}
	}
	return entity.Item{}, fmt.Errorf("no data")
}

func (m mockRepository) Update(userId int, item *entity.Item) error {
	for i, curItem := range m.data {
		if curItem.Id == item.Id {
			m.data[i] = *item
			return nil
		}
	}
	return fmt.Errorf("no data")
}

func NewMockRerository() *mockRepository {
	return &mockRepository{
		data: []entity.Item{
			{
				Id:       10,
				Name:     "sword",
				Rarity:   "rare",
				Category: "weapon",
				State:    entity.Equipped,
			},
			{
				Id:       6,
				Name:     "head",
				Rarity:   "legendary",
				Category: "armor",
				State:    entity.Equipped,
			},
		},
		userId: 1,
	}
}
