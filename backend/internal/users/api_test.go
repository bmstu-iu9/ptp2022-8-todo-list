package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
)

type ApiTestSuite struct {
	mux *httprouter.Router
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&ApiTestSuite{})
}

func (s *ApiTestSuite) SetUpTest(c *C) {
	s.mux = router.New()
	RegisterHandlers(s.mux, NewService(NewMockRerository()), log.New())
}

func (s *ApiTestSuite) TestPost(c *C) {
	makeRequest := func (body string) {
		s.writer = httptest.NewRecorder()
		bodyReader := strings.NewReader(body)
		request, _ := http.NewRequest("POST", "/users", bodyReader)
		s.mux.ServeHTTP(s.writer, request)
	}

	makeRequest(`{"email": "slava@example.com", "nickname": "slavarusvarrior", "password": "sDFHgjssndbfns123"}`)
	c.Check(s.writer.Code, Equals, http.StatusCreated)
	c.Check(s.writer.Header().Get("Location"), Equals, "https://ptp.starovoytovai.ru/api/v1/users/6")
	c.Check(s.writer.Body.Len(), Equals, 0)

	makeRequest(`"email": "slava@example.com", "nickname": "slavarusvarrior", "password": "sDFHgjssndbfns123"`)
	c.Check(s.writer.Code, Equals, http.StatusBadRequest)

	makeRequest(`{"email": "slava@example.com", "password": "sDFHgjssndbfns123"}`)
	c.Check(s.writer.Code, Equals, http.StatusBadRequest)

	makeRequest(`{"email": "slavaexample.com", "nickname": "slavarusvarrior", "password": "sDFHgjssndbfns123"}`)
	c.Check(s.writer.Code, Equals, http.StatusBadRequest)
}

func (s *ApiTestSuite) TestGet(c *C) {
	makeRequest := func (id string) {
		s.writer = httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/users/" + id, nil)
		s.mux.ServeHTTP(s.writer, request)
	}

	makeRequest("5")
	c.Check(s.writer.Code, Equals, http.StatusOK)
	got := User{}
	err := json.NewDecoder(s.writer.Body).Decode(&got)
	c.Check(err, Equals, nil)
	c.Check(got, DeepEquals, User {
		Id: 5,
		Email: "geogreck@example.com",
		Nickname: "geogreck",
	})

	makeRequest("6")
	c.Check(s.writer.Code, Equals, http.StatusNotFound)
}

func (s *ApiTestSuite) TestPatch(c *C) {
	makeRequest := func (id string, body string) {
		s.writer = httptest.NewRecorder()
		bodyReader := strings.NewReader(body)
		request, _ := http.NewRequest("PATCH", "/users/"+ id, bodyReader)
		s.mux.ServeHTTP(s.writer, request)
	}

	makeRequest("5", `{"email": "test@example.com", "currentPassword": "test123test"}`)
	c.Check(s.writer.Code, Equals, http.StatusOK)
	got := User{}
	err := json.NewDecoder(s.writer.Body).Decode(&got)
	c.Check(err, Equals, nil)
	c.Check(got, DeepEquals, User {
		Id: 5,
		Email: "test@example.com",
		Nickname: "geogreck",
	})

	makeRequest("6", `{"email": "test@example.com", "currentPassword": "test123test"}`)
	c.Check(s.writer.Code, Equals, http.StatusInternalServerError)

	makeRequest("5", `{"email": "test@example.com", "currentPassword": "test12test"}`)
	c.Check(s.writer.Code, Equals, http.StatusInternalServerError)

	makeRequest("5", `{"email": "testexample.com", "currentPassword": "test123test"}`)
	c.Check(s.writer.Code, Equals, http.StatusInternalServerError)
}

func (s *ApiTestSuite) TestDelete(c *C) {
	makeRequest := func (id string) {
		s.writer = httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/users/"+id, nil)
		s.mux.ServeHTTP(s.writer, request)
	}

	makeRequest("5")
	c.Check(s.writer.Code, Equals, http.StatusNoContent)
	// TODO Проверка наличия удаленного пользователя

	makeRequest("6")
	c.Check(s.writer.Code, Equals, http.StatusNotFound)
}
