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
		version:    bsVersion,
		baseURL:    bsBaseURL,
		httpClient: bs.httpClient,
	}
	c.Assert(bs, DeepEquals, expected)
}

func (s *MySuite) TestNewBSGetTokenWithoutAPIKey(c *C) {
	bs, err := NewBetaseriesClient("", "Dev050", "developer")
	c.Assert(err.Error(), Equals, "Veuillez spécifier une clé API.\n")
	c.Assert(bs, NotNil)
	expected := &BetaSeries{
		version:    bsVersion,
		baseURL:    bsBaseURL,
		httpClient: bs.httpClient,
	}
	c.Assert(bs, DeepEquals, expected)
}

func (s *MySuite) TestNewBSGetTokenWithAPIKey(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "Dev050", "developer")
	c.Assert(err, IsNil)
	c.Assert(bs, NotNil)
	token, err := bs.getToken()
	c.Assert(err, IsNil)
	c.Assert(len(token), Equals, 12)
}

func (s *MySuite) TestNewBSGetToken(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	c.Assert(bs, NotNil)
	_, err = bs.getToken()
	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrNoToken)
}

func makeClientAndAddShow(c *C) (*BetaSeries, string, int) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "Dev050", "developer")
	c.Assert(err, IsNil)
	_, err = bs.EpisodesList(0, 0)
	c.Assert(err, NotNil)
	// meaning null/nil return
	c.Assert(err.Error(), Equals, "")

	shows, err := bs.ShowsSearch(tvShowTest)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 1)

	// make sure the tv show is not in the user account first
	bs.ShowRemove(shows[0].ID)

	show, err := bs.ShowAdd(shows[0].ID)
	c.Assert(err, IsNil)
	c.Assert(show.InAccount, Equals, true)
	return bs, key, shows[0].ID
}
