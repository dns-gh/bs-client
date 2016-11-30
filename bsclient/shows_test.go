package bsclient

import (
	"os"

	. "gopkg.in/check.v1"
)

const (
	tvShowTest = "Breaking Bad"
)

var (
	err4001 = errorsAPI{
		Code: 4001,
		Text: "Aucune série trouvée.",
	}
)

func (s *MySuite) TestShowsSearch(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	shows, err := bs.ShowsSearch(tvShowTest)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 1)
	c.Assert(shows[0].ID, Equals, 481)
	c.Assert(shows[0].Title, Equals, tvShowTest)
	c.Assert(shows[0].Seasons, Equals, "5")
	c.Assert(shows[0].Episodes, Equals, "68")

	_, err = bs.ShowsSearch("TV Show doesn't exists")
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoShowsFound)
}

func (s *MySuite) TestShowsCharacters(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	shows, err := bs.ShowsSearch(tvShowTest)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 1)
	characters, err := bs.ShowsCharacters(shows[0].ID)
	c.Assert(err, IsNil)
	c.Assert(len(characters), Equals, 19)

	_, err = bs.ShowsCharacters(123456789)
	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, &errAPI{
		[]errorsAPI{err4001},
	})
}
