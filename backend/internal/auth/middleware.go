package auth

import (
	"fmt"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
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
		id, _ := strconv.Atoi(p.ByName("id"))
		isTokenValid := ValidateAccessToken(accessToken, id)
		if !isTokenValid {
			return fmt.Errorf("%w: %v", errors.ErrUnauthorized, "invalid access token")
		}
		return next(w, r, p)
	}
}
