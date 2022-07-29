package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
)

type PingTestSuite struct{
	mux *httprouter.Router
	writer *httptest.ResponseRecorder
	logger log.Logger
}

func init() {
	Suite(&PingTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PingTestSuite) SetUpSuite(c *C) {
	s.mux = httprouter.New()
	s.logger = log.NewForTest()
	RegisterHandlers(s.mux, s.logger)
	s.writer = httptest.NewRecorder()
}

func (s *PingTestSuite) TestPing(c *C) {
	request, _ := http.NewRequest("GET", "/ping", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 418)
	c.Check(s.writer.Body.Len(), Equals, 0)
}
