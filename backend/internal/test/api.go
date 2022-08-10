package test

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ApiTestCases represents the data needed to describe an API endpoint test cases
type ApiTestCase struct {
	Name           string
	Method         string
	Url            string
	Body           string
	Cookie         http.Cookie
	CookieRequired bool
	Header         http.Header
	WantCookie     http.Cookie
	WantCode       int
	WantBody       string
	WantHeader     http.Header
}

// Endpoint tests the API endpoint using given test cases.
func Endpoint(t *testing.T, testCases []ApiTestCase, mux *httprouter.Router) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			request, _ := http.NewRequest(tc.Method, tc.Url, strings.NewReader(tc.Body))
			writer := httptest.NewRecorder()
			if tc.CookieRequired {
				http.SetCookie(writer, &tc.Cookie)
				request.Header = http.Header{"Cookie": writer.Result().Header["Set-Cookie"]}
			}
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
