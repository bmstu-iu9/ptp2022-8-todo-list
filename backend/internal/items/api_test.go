package items

import (
	"encoding/json"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
	RegisterHandlers(s.mux, NewService(NewMockRerository()), log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile))
}

func Test(t *testing.T) { TestingT(t) }

func (s *ApiTestSuite) TestGetAll(c *C) {
	makeRequest := func(userId string) {
		s.writer = httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users/:"+userId+"/items", nil)
		s.mux.ServeHTTP(s.writer, req)
	}
	makeRequest("1") //any userId
	c.Check(s.writer.Code, Equals, http.StatusOK)
	got := []Item{}
	fmt.Println(s.writer.Body)
	err := json.NewDecoder(s.writer.Body).Decode(&got)
	fmt.Println(got)
	c.Check(err, Equals, nil)
	c.Check(got, DeepEquals, []Item{
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

type mockRepository struct {
	items  []Item
	userId int
}

func (m mockRepository) GetAll() ([]Item, error) {
	return m.items, nil
}

func (m mockRepository) GetOne(user *entity.User, item *entity.Item) (entity.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockRepository) Modify(user *entity.User, item *entity.Item) error {
	//TODO implement me
	panic("implement me")
}

func NewMockRerository() *mockRepository {
	return &mockRepository{
		items: []Item{
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
