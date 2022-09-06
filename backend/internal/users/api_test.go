package users

import (
	"bytes"
	"encoding/json"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/auth"
	"net/http"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
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
	mux = router.New(logger)
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

func GenerateBearerAccessToken(email entity.Email, userId int) string {
	tokens, _ := auth.GenerateTokens(email, userId)
	return "Bearer " + tokens.AccessToken
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
	forbidden := toJson(errors.Problem{Title: "Forbidden", Status: http.StatusForbidden, Detail: "Wrong login or password"})

	tests := []test.ApiTestCase{
		{Name: "create OK", Method: "POST", Url: "/users",
			Body:     `{"email": "slava@example.com", "nickname": "slavarusvarrior", "password": "sDFHgjssndbfns123"}`,
			WantCode: http.StatusCreated},
		{Name: "create verify", Method: "GET", Url: "/users/2",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("slava@example.com", 2)}},
			WantBody: toJson(entity.UserDto{Id: 2, Email: "slava@example.com", Nickname: "slavarusvarrior"}),
			WantCode: http.StatusOK},
		{Name: "create input error", Method: "POST", Url: "/users",
			Body: `"email": "slava@example.com", "nickname": "slavarusvarrior", "password": "sDFHgjssndbfns123"`,
			WantBody: toJson(errors.Problem{Title: "Bad request", Status: http.StatusBadRequest, Detail: "Bad request body"}),
			WantCode: http.StatusBadRequest},
		{Name: "create input error", Method: "POST", Url: "/users",
			Body: `{"email": "slava@example.com", "password": "sDFHgjssndbfns123"}`,
			WantBody: toJson(errors.Problem{Title: "Bad request", Status: http.StatusBadRequest, Detail: "Request body parameters validation failed"}),
			WantCode: http.StatusBadRequest},
		{Name: "get OK", Method: "GET", Url: "/users/1",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			WantCode: http.StatusOK, WantBody: toJson(entity.UserDto{Id: 1, Email: "geogreck@example.com", Nickname: "geogreck"})},
		{Name: "get id error", Method: "GET", Url: "/users/33",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("33@example.com", 33)}},
			WantBody: notFound,
			WantCode: http.StatusNotFound},
		{Name: "get auth error", Method: "GET", Url: "/users/1",
			WantBody: unauthorized,
			WantCode: http.StatusUnauthorized},
		{Name: "modify OK", Method: "PATCH", Url: "/users/1",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			Body:     `{"email": "test@example.com", "currentPassword": "Test123Test"}`,
			WantCode: http.StatusOK, WantBody: toJson(entity.UserDto{Id: 1, Email: "test@example.com", Nickname: "geogreck"})},
		{Name: "modify verify", Method: "GET", Url: "/users/1",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			WantCode: http.StatusOK, WantBody: toJson(entity.UserDto{Id: 1, Email: "test@example.com", Nickname: "geogreck"})},
		{Name: "modify id error", Method: "PATCH", Url: "/users/33",
			Body:     `{"email": "test@example.com", "currentPassword": "Test123Test"}`,
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("33@example.com", 33)}},
			WantBody: notFound,
			WantCode: http.StatusNotFound},
		{Name: "modify current password error", Method: "PATCH", Url: "/users/1",
			Body:     `{"email": "test@example.com", "currentPassword": "test12test"}`,
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			WantBody: forbidden,
			WantCode: http.StatusForbidden},
		{Name: "modify verify", Method: "GET", Url: "/users/1",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			WantCode: http.StatusOK, WantBody: toJson(entity.UserDto{Id: 1, Email: "test@example.com", Nickname: "geogreck"})},
		{Name: "modify input error", Method: "PATCH", Url: "/users/1",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			Body:     `{"email": "testexample.com", "currentPassword": "Test123Test"}`,
			WantBody: toJson(errors.Problem{Title: "Bad request", Status: http.StatusBadRequest, Detail: "Request body parameters validation failed"}),
			WantCode: http.StatusBadRequest},
		{Name: "delete OK", Method: "DELETE", Url: "/users/1",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			WantCode: http.StatusNoContent},
		{Name: "delete verify", Method: "DELETE", Url: "/users/1",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("geogreck@example.com", 1)}},
			WantBody: notFound,
			WantCode: http.StatusNotFound},
		{Name: "delete id error", Method: "DELETE", Url: "/users/33",
			Header:   http.Header{"Authorization": []string{GenerateBearerAccessToken("33@example.com", 33)}},
			WantBody: notFound,
			WantCode: http.StatusNotFound},
	}

	test.Endpoint(t, tests, mux)
}
