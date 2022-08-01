package items

import (
	"encoding/json"
	"errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
)

type ApiTestSuite struct {
	mux    *httprouter.Router
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&ApiTestSuite{})
}

func (s *ApiTestSuite) SetUpTest(c *C) {
	s.mux = httprouter.New()
	RegisterHandlers(s.mux, NewService(NewMockRerository()), log.New())
}

func Test(t *testing.T) { TestingT(t) }

func (s *ApiTestSuite) TestGetAll(c *C) {
	makeRequest := func(userId string) {
		s.writer = httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users/"+userId+"/items", nil)
		s.mux.ServeHTTP(s.writer, req)
	}
	makeRequest("1") //any userId
	c.Check(s.writer.Code, Equals, http.StatusOK)
	got := []entity.Item{}
	err := json.NewDecoder(s.writer.Body).Decode(&got)
	c.Check(err, Equals, nil)
	c.Check(got, DeepEquals, []entity.Item{
		{
			ItemId:   10,
			ItemName: "sword",
		},
		{
			ItemId:   6,
			ItemName: "head",
		},
	})
}

func (s *ApiTestSuite) TestGetOne(c *C) {
	makeRequest := func(userId, itemId string) {
		s.writer = httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users/"+userId+"/items/"+itemId, nil)
		s.mux.ServeHTTP(s.writer, req)
	}
	makeRequest("1", "10")
	c.Check(s.writer.Code, Equals, http.StatusOK)
	got := entity.Item{}
	err := json.NewDecoder(s.writer.Body).Decode(&got)
	c.Check(err, Equals, nil)
	c.Check(got, DeepEquals, entity.Item{
		ItemId:   10,
		ItemName: "sword",
	})
	makeRequest("4", "10")
	c.Check(s.writer.Code, Equals, http.StatusNotFound)
	makeRequest("1", "5")
	c.Check(s.writer.Code, Equals, http.StatusNotFound)
}

func (s *ApiTestSuite) TestPatch(c *C) {
	makeRequest := func(userId, itemId, body string) {
		s.writer = httptest.NewRecorder()
		bodyReader := strings.NewReader(body)
		req, _ := http.NewRequest("PATCH", "/users/"+userId+"/items/"+itemId, bodyReader)
		s.mux.ServeHTTP(s.writer, req)
	}

	makeRequest("1", "10", `{"item_state":"equipped"}`)
	c.Check(s.writer.Code, Equals, http.StatusOK)
	got := entity.Item{}
	err := json.NewDecoder(s.writer.Body).Decode(&got)
	c.Check(err, IsNil)
	c.Check(got, DeepEquals, entity.Item{
		ItemId:    10,
		ItemName:  "sword",
		ItemState: entity.Equipped,
	})
	makeRequest("1", "2", `{"ItemName": "test"}`)
	c.Check(s.writer.Code, Equals, http.StatusInternalServerError)
	makeRequest("10", "10", `{"ItemName": "test"}`)
	c.Check(s.writer.Code, Equals, http.StatusInternalServerError)
}

type mockRepository struct {
	data   []entity.Item
	userId int
}

func (m mockRepository) GetAll() ([]entity.Item, error) {
	return m.data, nil
}

func (m mockRepository) GetOne(userId, itemId int) (entity.Item, error) {
	if userId != m.userId {
		return entity.Item{}, errors.New("No user")
	}
	for _, item := range m.data {
		if item.ItemId == itemId {
			return item, nil
		}
	}
	return entity.Item{}, errors.New("No data")
}

func (m mockRepository) Update(userId int, item *entity.Item) error {
	for i, curItem := range m.data {
		if curItem.ItemId == item.ItemId {
			m.data[i] = *item
			return nil
		}
	}
	return errors.New("No data")
}

func NewMockRerository() *mockRepository {
	return &mockRepository{
		data: []entity.Item{
			{
				ItemId:   10,
				ItemName: "sword",
			},
			{
				ItemId:   6,
				ItemName: "head",
			},
		},
		userId: 1,
	}
}
