package bsclient

import (
	"os"
	"strings"

	. "gopkg.in/check.v1"
)

var (
	err0 = errorsAPI{
		Code: 0,
		Text: "Aucun utilisateur sélectionné.",
	}
)

func checkEpisode(c *C, err error, episode *Episode) {
	c.Assert(err, IsNil)
	c.Assert(len(episode.Date), Equals, 10)
	c.Assert(strings.Contains(episode.Code, "S"), Equals, true)
	c.Assert(strings.Contains(episode.Code, "E"), Equals, true)
}

func (s *MySuite) TestPlanningGeneral(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	episodes, err := bs.PlanningGeneral("now", 1, 1)
	if len(episodes) > 0 {
		checkEpisode(c, err, &episodes[0])
	} else {
		c.Assert(err, NotNil)
	}

	episodes, err = bs.PlanningGeneral("1000-01-01", 1, 1)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoEpisodesFound)
}

func (s *MySuite) TestPlanningIncoming(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	episodes, err := bs.PlanningIncoming()
	if len(episodes) > 0 {
		checkEpisode(c, err, &episodes[0])
	} else {
		c.Assert(err, NotNil)
		c.Assert(err, Equals, errNoEpisodesFound)
	}
}

func (s *MySuite) TestPlanningMember(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	episodes, err := bs.PlanningMember(0, false, "")
	if len(episodes) > 0 {
		checkEpisode(c, err, &episodes[0])
	} else {
		c.Assert(err, NotNil)
		c.Assert(err, DeepEquals, &errAPI{
			[]errorsAPI{err0},
		})
	}

	episodes, err = bs.PlanningMember(-1, false, "")
	if len(episodes) > 0 {
		checkEpisode(c, err, &episodes[0])
	} else {
		c.Assert(err, NotNil)
		c.Assert(err, DeepEquals, &errAPI{
			[]errorsAPI{err0},
		})
	}

	episodes, err = bs.PlanningMember(-1, false, "1000-01")
	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, &errAPI{
		[]errorsAPI{err0},
	})

	episodes, err = bs.PlanningMember(-1, false, "Wrong format")
	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, &errAPI{
		[]errorsAPI{err0},
	})
}

func (s *MySuite) TestPlanningMemberWithCredentials(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "Dev050", "developer")
	c.Assert(err, IsNil)
	episodes, err := bs.PlanningMember(0, false, "")
	if len(episodes) > 0 {
		checkEpisode(c, err, &episodes[0])
	} else {
		c.Assert(err, NotNil)
		c.Assert(err, Equals, errNoEpisodesFound)
	}

	episodes, err = bs.PlanningMember(-1, false, "")
	if len(episodes) > 0 {
		checkEpisode(c, err, &episodes[0])
	} else {
		c.Assert(err, NotNil)
		c.Assert(err, Equals, errNoEpisodesFound)
	}

	episodes, err = bs.PlanningMember(-1, false, "1000-01")
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoEpisodesFound)

	episodes, err = bs.PlanningMember(-1, false, "Wrong format")
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "DateTime::__construct(): Failed to parse time string (Wrong format) at position 0 (W): The timezone could not be found in the database\n")

	episodes, err = bs.PlanningMember(-1, false, "now")
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoEpisodesFound)
}
