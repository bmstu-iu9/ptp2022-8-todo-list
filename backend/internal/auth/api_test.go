package auth

import (
	"bytes"
	"encoding/json"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
	"time"
)

var (
	mux    *httprouter.Router
	logger log.Logger
	s      service
	tokens Token
)

func init() {
	mux = router.New()
	logger = log.New()
	tokens, _ = generateTokens("slava@example.com")
	s = service{&mockRepository{
		users: []entity.User{
			{
				Id:       0,
				Email:    "slava@example.com",
				Nickname: "slavaruswarrior",
				Password: "3dfff1ca8a9696f67616a2b8abd1bce3", //wasdqwertytest
			},
			{
				Id:       5,
				Email:    "geogreck@example.com",
				Nickname: "geogreck",
				Password: "test123test",
			},
		},
		tokens: []DbToken{
			{
				userId:       5,
				refreshToken: "token",
			},
			{
				userId:       0,
				refreshToken: tokens.RefreshToken,
			},
		},
	}}
	RegisterHandlers(mux, s, logger)
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
	badRequest := toJson(errors.Problem{Title: "Bad request", Status: http.StatusBadRequest})
	unauthorized := toJson(errors.Problem{Title: "Unauthorized", Status: http.StatusUnauthorized})
	tests := []test.ApiTestCase{
		{
			Name: "login OK", Method: "POST", Url: "/login",
			Body:     toJson(LoginUserRequest{Email: "slava@example.com", Password: "wasdqwertytest"}),
			WantCode: http.StatusOK,
			WantBody: toJson(UserData{
				User: entity.UserDto{
					Id:       0,
					Email:    "slava@example.com",
					Nickname: "slavaruswarrior",
				},
				Tokens: tokens,
			}),
		},
		{
			Name: "login FAIL: wrong pass", Method: "POST", Url: "/login",
			Body:     toJson(LoginUserRequest{Email: "slava@example.com", Password: "pass"}),
			WantCode: http.StatusBadRequest,
			WantBody: badRequest,
		},
		{
			Name: "login FAIL: no email in db", Method: "POST", Url: "/login",
			Body:     toJson(LoginUserRequest{Email: "sarahdeep@example.com", Password: "wasdqwertytest"}),
			WantCode: http.StatusBadRequest,
			WantBody: badRequest,
		},
		{
			Name: "logout OK", Method: "POST", Url: "/logout",
			CookieRequired: true,
			Cookie: http.Cookie{
				Name:     "refreshToken",
				Value:    "token",
				Expires:  time.Now().Add(30 * 24 * time.Hour),
				HttpOnly: true,
			},
			WantCode: http.StatusOK,
		},
		{
			Name: "logout Fail: no token in db", Method: "POST", Url: "/logout",
			CookieRequired: true,
			Cookie: http.Cookie{
				Name:     "refreshToken",
				Value:    "token",
				Expires:  time.Now().Add(30 * 24 * time.Hour),
				HttpOnly: true,
			},
			WantCode: http.StatusBadRequest,
			WantBody: badRequest,
		},
		{
			Name: "logout Fail: unauthorized", Method: "POST", Url: "/logout",
			CookieRequired: true,
			Cookie: http.Cookie{
				Name:     "fakeToken",
				Value:    "token",
				Expires:  time.Now().Add(30 * 24 * time.Hour),
				HttpOnly: true,
			},
			WantCode: http.StatusUnauthorized,
			WantBody: unauthorized,
		},
		{
			Name: "refresh OK", Method: "GET", Url: "/refresh",
			CookieRequired: true,
			Cookie: http.Cookie{
				Name:     "refreshToken",
				Value:    tokens.RefreshToken,
				Expires:  time.Now().Add(30 * 24 * time.Hour),
				HttpOnly: true,
			},
			WantBody: toJson(UserData{
				User: entity.UserDto{
					Id:       0,
					Email:    "slava@example.com",
					Nickname: "slavaruswarrior",
				},
				Tokens: tokens,
			}),
			WantCode: http.StatusOK,
		},
		{
			Name: "refresh Fail: wrong token", Method: "GET", Url: "/refresh",
			CookieRequired: true,
			Cookie: http.Cookie{
				Name:     "refreshToken",
				Value:    "badToken",
				Expires:  time.Now().Add(30 * 24 * time.Hour),
				HttpOnly: true,
			},
			WantCode: http.StatusUnauthorized,
			WantBody: unauthorized,
		},
	}
	test.Endpoint(t, tests, mux)
}
