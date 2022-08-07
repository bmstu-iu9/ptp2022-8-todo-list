package auth

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func RegisterHandlers(mux *httprouter.Router, service Service, logger *log.Logger) {
	res := resource{service, logger}

	mux.POST("/login", res.handleLog)
	mux.POST("/logout", res.handleLogOut)
	mux.GET("/refresh", res.handleRefresh)
}

type resource struct {
	service Service
	logger  *log.Logger
}

func (res *resource) handleLog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := LoginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	user, tokens, err := res.service.Login(data.Email, data.Password)
	if err != nil {
		res.logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	})
	err = json.NewEncoder(w).Encode(user)
	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		res.logger.Println(err)
		return
	}
}

func (res *resource) handleLogOut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		res.logger.Println(err)
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	refreshToken := cookie.Value
	err = res.service.Logout(refreshToken)
	if err != nil {
		res.logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshToken",
		Value:   "",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
}

func (res *resource) handleRefresh(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		res.logger.Println(err)
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	refreshToken := cookie.Value
	user, tokens, err := res.service.Refresh(refreshToken)
	if err != nil {
		res.logger.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	})
	err = json.NewEncoder(w).Encode(user)
	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		res.logger.Println(err)
		return
	}
}
