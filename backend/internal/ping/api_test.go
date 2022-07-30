package ping

import (
	"net/http"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

func TestApi(t *testing.T) {
	mux := router.New()
	logger := log.New()
	RegisterHandlers(mux, logger)

	tests := []test.ApiTestCase{
		{Name: "Ping OK", Method: "GET", Url: "/ping", Body: "",
			WantCode: http.StatusTeapot},
		{Name: "Ping non-empty body", Method: "GET", Url: "/ping", Body: "12345",
			WantCode: http.StatusTeapot},
		{Name: "Ping wrong method", Method: "POST", Url: "/ping", Body: "{}",
			WantCode: http.StatusMethodNotAllowed, WantBody: "Method Not Allowed\n"},
	}

	test.Endpoint(t, tests, mux)
}
