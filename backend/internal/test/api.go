package test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// ApiTestCases represents the data needed to describe an API endpoint test cases
type ApiTestCases map[string]struct {
	Method     string
	Url        string
	Body       string
	Header     http.Header
	WantCode int
	WantBody   string
	WantHeader http.Header
}

// Endpoint tests the API endpoint using given test cases.
func Endpoint(t *testing.T, testCases ApiTestCases, mux *httprouter.Router) {
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			request, _ := http.NewRequest(tc.Method, tc.Url, strings.NewReader(tc.Body))
			writer := httptest.NewRecorder()
			mux.ServeHTTP(writer, request)

			gotCode := writer.Code
			gotBody := writer.Body.String()
			gotHeader := writer.Header()

			if tc.WantCode != gotCode {
				t.Fatalf("expected: %#v, got: %#v", tc.WantCode, gotCode)
			}
			if !reflect.DeepEqual(tc.WantBody, gotBody) {
				t.Fatalf("expected: %#v, got: %#v", tc.WantBody, gotBody)
			}
			for name, wantValue := range tc.WantHeader {
				gotValue := gotHeader.Values(name)
				if !reflect.DeepEqual(wantValue, gotValue) {
					t.Fatalf("expected header %s value: %#v, got: %#v", name, wantValue, gotValue)
				}
			}
		})
	}
}
