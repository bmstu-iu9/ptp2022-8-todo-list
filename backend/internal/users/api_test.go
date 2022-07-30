package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
	"github.com/julienschmidt/httprouter"
)

var (
	mux    *httprouter.Router
	logger log.Logger
)

func init() {
	mux = router.New()
	logger = log.New()
	service := NewService(&mockRepository{
		id: 2,
		items: []entity.User{{
			Id:       1,
			Email:    "geogreck@example.com",
			Nickname: "geogreck",
			Password: "Test123Test",
		}},
	})
	RegisterHandlers(mux, service, logger)
}

func TestApi(t *testing.T) {
	toJson := func(user User) string {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(user)
		if err != nil {
			panic(err)
		}
		return buf.String()
	}

	tests := []test.ApiTestCase{
		{Name: "create OK", Method: "POST", Url: "/users",
			Body:     `{"email": "slava@example.com", "nickname": "slavarusvarrior", "password": "sDFHgjssndbfns123"}`,
			WantCode: http.StatusCreated,
			WantHeader: http.Header{"Location": {"https://ptp.starovoytovai.ru/api/v1/users/2"}}},
		{Name: "create verify", Method: "GET", Url: "/users/2",
			WantBody: toJson(User{Id: 2, Email: "slava@example.com", Nickname: "slavarusvarrior"}),
			WantCode: http.StatusOK},
		{Name: "create input error", Method: "POST", Url: "/users",
			Body: `"email": "slava@example.com", "nickname": "slavarusvarrior", "password": "sDFHgjssndbfns123"`,
			WantCode: http.StatusBadRequest},
		{Name: "Create input error", Method: "POST", Url: "/users",
			Body: `{"email": "slava@example.com", "password": "sDFHgjssndbfns123"}`,
			WantCode: http.StatusBadRequest},
		{Name: "get OK", Method: "GET", Url: "/users/1",
			WantCode: http.StatusOK, WantBody: toJson(User{Id: 1, Email: "geogreck@example.com", Nickname: "geogreck"})},
		{Name: "get id error", Method: "GET", Url: "/users/33",
			WantCode: http.StatusNotFound},
		{Name: "modify OK", Method: "PATCH", Url: "/users/1",
			Body:     `{"email": "test@example.com", "currentPassword": "Test123Test"}`,
			WantCode: http.StatusOK, WantBody: toJson(User{Id: 1, Email: "test@example.com", Nickname: "geogreck"})},
		{Name: "modify verify", Method: "GET", Url: "/users/1",
			WantCode: http.StatusOK, WantBody: toJson(User{Id: 1, Email: "test@example.com", Nickname: "geogreck"})},
		{Name: "modify id error", Method: "PATCH", Url: "/users/33",
			WantCode: http.StatusNotFound},
		{Name: "modify current password error", Method: "PATCH", Url: "/users/1",
			Body:     `{"email": "test@example.com", "currentPassword": "test12test"}`,
			WantCode: http.StatusForbidden},
		{Name: "modify verify", Method: "GET", Url: "/users/1",
			WantCode: http.StatusOK, WantBody: toJson(User{Id: 1, Email: "test@example.com", Nickname: "geogreck"})},
		{Name: "modify input error", Method: "PATCH", Url: "/users/1",
			Body:     `{"email": "testexample.com", "currentPassword": "Test123Test"}`,
			WantCode: http.StatusBadRequest},
		{Name: "delete OK", Method: "DELETE", Url: "/users/1",
			WantCode: http.StatusNoContent},
		{Name: "delete verify", Method: "DELETE", Url: "/users/1",
			WantCode: http.StatusNotFound},
		{Name: "delete id error", Method: "DELETE", Url: "/users/33",
			WantCode: http.StatusNotFound},
	}

	test.Endpoint(t, tests, mux)
}
