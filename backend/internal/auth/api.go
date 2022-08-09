package auth

import (
	"encoding/json"
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
	userData, err := res.service.Login(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    userData.Tokens.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
	})
	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		return err
	}
	return nil
}

func (res *resource) handleLogOut(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return err
	}
	refreshToken := cookie.Value
	err = res.service.Logout(refreshToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshToken",
		Value:   "",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
	return nil
}

func (res *resource) handleRefresh(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return err
	}
	refreshToken := cookie.Value
	userData, err := res.service.Refresh(refreshToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    userData.Tokens.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
	})
	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		return err
	}
	return nil
}
