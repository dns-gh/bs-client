package bsclient

import (
	"os"

	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestPlanningGeneral(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	episodes, err := bs.PlanningGeneral("now", 1, 1)
	c.Assert(err, IsNil)
	if len(episodes) > 0 {
		c.Assert(len(episodes[0].Date), Equals, 10)
		c.Assert(strings.Contains(episodes[0].Code, "S"), Equals, true)
		c.Assert(strings.Contains(episodes[0].Code, "E"), Equals, true)
		c.Assert(len(episodes[0].Date), Equals, 10)
	}

	episodes, err = bs.PlanningGeneral("1000-01-01", 1, 1)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoEpisodesFound)
}
