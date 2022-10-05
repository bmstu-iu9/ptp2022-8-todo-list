package ping

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

func TestApi(t *testing.T) {
	toJson := func(data interface{}) string {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(data)
		if err != nil {
			panic(err)
		}
		return buf.String()
	}

	logger := log.New()
	mux := router.New(logger)
	RegisterHandlers(mux, logger)

	tests := []test.ApiTestCase{
		{Name: "Ping OK", Method: "GET", Url: "/ping", Body: "",
			WantCode: http.StatusTeapot},
		{Name: "Ping non-empty body", Method: "GET", Url: "/ping", Body: "12345",
			WantCode: http.StatusTeapot},
		{Name: "Ping wrong method", Method: "POST", Url: "/ping", Body: "{}",
			WantCode: http.StatusMethodNotAllowed, WantBody: toJson(errors.Problem{
				Title:  "Method not allowed",
				Status: http.StatusMethodNotAllowed,
			})},
	}

	test.Endpoint(t, tests, mux)
}
