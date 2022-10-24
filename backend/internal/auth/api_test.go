package auth

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type AuthSuite struct{}

var _ = Suite(&AuthSuite{})

func (s *AuthSuite) TestReturns200(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

}
