package bsclient

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

// go test ...bsclient -gocheck.vv -test.v -gocheck.f TestNAME
func (s *MySuite) TestNewBS(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs := NewBetaseriesClient(key, "", "")
	c.Assert(bs, NotNil)
	expected := &BetaSeries{
		version: bsVersion,
		baseURL: bsBaseURL,
		key:     key,
	}
	c.Assert(bs, DeepEquals, expected)
}
