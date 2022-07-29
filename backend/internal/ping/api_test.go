package ping

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
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
	tests := map[string]struct {
		method   string
		body     string
		wantCode int
		wantBody string
	}{
		"OK":             {method: "GET", body: "", wantCode: http.StatusTeapot, wantBody: ""},
		"Non-empty body": {method: "GET", body: "12345", wantCode: http.StatusTeapot, wantBody: ""},
		"Wrong method":   {method: "POST", body: "{}", wantCode: http.StatusMethodNotAllowed, wantBody: "Method Not Allowed\n"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			request, _ := http.NewRequest(tc.method, "/ping", strings.NewReader(tc.body))
			writer := httptest.NewRecorder()
			mux.ServeHTTP(writer, request)

			gotCode := writer.Code
			gotBody := writer.Body.String()

			if tc.wantCode != gotCode {
				t.Fatalf("expected: %#v, got: %#v", tc.wantCode, gotCode)
			}
			if !reflect.DeepEqual(tc.wantBody, gotBody) {
				t.Fatalf("expected: %#v, got: %#v", tc.wantBody, gotBody)
			}
		})
	}
}
