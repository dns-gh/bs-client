package bsclient

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

// export BS_API_KEY=YOUR_API_KEY && go test ...bsclient -gocheck.vv -test.v -gocheck.f TestNAME
func (s *MySuite) TestNewBS(c *C) {
	bs, err := NewBetaseriesClient("", "", "")
	c.Assert(err, IsNil)
	expected := &BetaSeries{
		version: bsVersion,
		baseURL: bsBaseURL,
	}
	c.Assert(bs, DeepEquals, expected)
}

func (s *MySuite) TestNewBSGetTokenWithoutAPIKey(c *C) {
	bs, err := NewBetaseriesClient("", "Dev050", "developer")
	c.Assert(err.Error(), Equals, "Veuillez spécifier une clé API.\n")
	c.Assert(bs, NotNil)
	expected := &BetaSeries{
		version: bsVersion,
		baseURL: bsBaseURL,
	}
	c.Assert(bs, DeepEquals, expected)
}

func (s *MySuite) TestNewBSGetToken(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "Dev050", "developer")
	c.Assert(err, IsNil)
	c.Assert(bs, NotNil)
	token, err := bs.getToken()
	c.Assert(err, IsNil)
	c.Assert(len(token), Equals, 12)
}
