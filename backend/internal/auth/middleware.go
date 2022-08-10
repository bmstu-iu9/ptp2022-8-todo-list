package auth

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func AuthCheck(next errors.Handler) errors.Handler {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			return errors.ErrUnauthorized
		}
		accessToken := strings.Split(authHeader, " ")[1]
		if accessToken == "" {
			return errors.ErrUnauthorized
		}
		isTokenValid := ValidateAccessToken(accessToken)
		if !isTokenValid {
			return errors.ErrUnauthorized
		}
		return next(w, r, p)
	}
}
