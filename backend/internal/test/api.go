package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// ApiTestCases represents the data needed to describe an API endpoint test cases
type ApiTestCase struct {
	Name       string
	Method     string
	Url        string
	Body       string
	Header     http.Header
	WantCode   int
	WantBody   string
	WantHeader http.Header
}

// Endpoint tests the API endpoint using given test cases.
func Endpoint(t *testing.T, testCases []ApiTestCase, mux *httprouter.Router) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			request, _ := http.NewRequest(tc.Method, tc.Url, strings.NewReader(tc.Body))
			writer := httptest.NewRecorder()
			mux.ServeHTTP(writer, request)

			gotCode := writer.Code
			gotBody := writer.Body.String()
			gotHeader := writer.Header()

			if tc.WantCode != gotCode {
				t.Fatalf("expected: %#v, got: %#v", tc.WantCode, gotCode)
			}
			DeepEqual(t, tc.WantBody, gotBody)
			for name, wantValue := range tc.WantHeader {
				gotValue := gotHeader.Values(name)
				DeepEqual(t, wantValue, gotValue)
			}
		})
	}
}
