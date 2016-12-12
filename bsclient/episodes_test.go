package bsclient

import (
	"os"

	. "gopkg.in/check.v1"
)

var (
	err2001 = errorsAPI{
		Code: 2001,
		Text: "Token invalide.",
	}
)

func (s *MySuite) TestEpisodesList(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	_, err = bs.EpisodesList(0, 0)
	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, &errAPI{
		[]errorsAPI{err2001},
	})

	bs, err = NewBetaseriesClient(key, "Dev050", "developer")
	c.Assert(err, IsNil)
	_, err = bs.EpisodesList(0, 0)
	c.Assert(err, NotNil)
	// meaning null/nil return
	c.Assert(err.Error(), Equals, "")

	_, err = bs.EpisodesList(-1, -1)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoShowsFound)
}
