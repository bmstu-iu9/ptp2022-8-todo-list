package auth

import (
	"encoding/json"
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/accesslog"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func RegisterHandlers(mux *httprouter.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	mux.POST("/login", accesslog.Log(errors.Handle(res.handleLog, logger), logger))
	mux.POST("/logout", accesslog.Log(errors.Handle(res.handleLogOut, logger), logger))
	mux.GET("/refresh", accesslog.Log(errors.Handle(res.handleRefresh, logger), logger))
}

type resource struct {
	service Service
	logger  log.Logger
}

func (res *resource) handleLog(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	data := LoginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return wrapDecode(err)
	}
	userData, err := res.service.Login(data)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrAuthentication, err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    userData.Tokens.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
	})
	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		return wrapEncode(err)
	}
	return nil
}

func (res *resource) handleLogOut(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrAuthentication, err)
	}
	refreshToken := cookie.Value
	err = res.service.Logout(refreshToken)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshToken",
		Value:   "",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})
	return nil
}

func (res *resource) handleRefresh(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrAuthentication, err)
	}
	refreshToken := cookie.Value
	userData, err := res.service.Refresh(refreshToken)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrAuthentication, err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    userData.Tokens.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
	})
	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		return wrapEncode(err)
	}
	return nil
}

func wrapDecode(err error) error {
	return fmt.Errorf("%w: %v", errors.ErrBodyDecode, err)
}

func wrapEncode(err error) error {
	return fmt.Errorf("%w: %v", errors.ErrBodyEncode, err)
}
