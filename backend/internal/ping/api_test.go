package ping

import (
	"net/http"
	"testing"

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
	mux = router.New()
	logger = log.New()
	RegisterHandlers(mux, logger)
}

func TestPing(t *testing.T) {
	tests := test.ApiTestCases{
		"OK": {Method: "GET", Url: "/ping", Body: "",
			WantCode: http.StatusTeapot},
		"Non-empty body": {Method: "GET", Url: "/ping", Body: "12345",
			WantCode: http.StatusTeapot},
		"Wrong method": {Method: "POST", Url: "/ping", Body: "{}",
			WantCode: http.StatusMethodNotAllowed, WantBody: "Method Not Allowed\n"},
	}

	test.Endpoint(t, tests, mux)
}
