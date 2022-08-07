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
	err = json.NewEncoder(w).Encode(user)
	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		res.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshToken",
		Value:   tokens.RefreshToken,
		Expires: time.Now().Add(30 * 24 * time.Hour),
	})
}

func (res *resource) handleLogOut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (res *resource) handleRefresh(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
