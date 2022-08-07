package auth

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func AuthCheck(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		accessToken := strings.Split(authHeader, " ")[1]
		if accessToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		isTokenValid := validateAccessToken(accessToken)
		if !isTokenValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r, p)
	}
}
