package bsclient

import (
	"os"
	"strings"

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

func (s *MySuite) TestShowsRandom(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	shows, err := bs.ShowsRandom(1, false)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 1)
	c.Assert(len(shows[0].Language) > 0, Equals, true)

	shows, err = bs.ShowsRandom(0, false)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoShowsFound)

	shows, err = bs.ShowsRandom(1, true)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 1)
	c.Assert(len(shows[0].Language), Equals, 0)
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

func (s *MySuite) TestShowsList(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	shows, err := bs.ShowsList("", "", -1, 100)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 100)
	c.Assert(shows[0].ID, Equals, 425)

	shows, err = bs.ShowsList("", "", 1, 100)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 100)
	c.Assert(shows[0].ID, Equals, 481)

	// timestamp to 01-01-3000
	shows, err = bs.ShowsList("32503680000", "", 1, 100)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoShowsFound)

	// timestamp to 01-01-2016
	shows, err = bs.ShowsList("1451606400", "", 1, 100)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 100)

	shows, err = bs.ShowsList("-wrong-", "", 1, 100)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 100)

	shows, err = bs.ShowsList("1451606400", "test", -1, 10)
	c.Assert(err, IsNil)
	c.Assert(len(shows), Equals, 1)
	c.Assert(shows[0].ID, Equals, 13842)
}

func (s *MySuite) TestShowsUpdate(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "Dev050", "developer")
	show, err := bs.ShowAdd(0)
	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, &errAPI{
		[]errorsAPI{err4001},
	})

	bs, _, id := makeClientAndAddShow(c)

	show, err = bs.ShowArchive(id)
	c.Assert(err, IsNil)
	c.Assert(show.InAccount, Equals, true)
	c.Assert(show.User.Archived, Equals, true)

	show, err = bs.ShowNotArchive(id)
	c.Assert(err, IsNil)
	c.Assert(show.InAccount, Equals, true)
	c.Assert(show.User.Archived, Equals, false)

	show, err = bs.ShowRemove(id)
	c.Assert(err, IsNil)
	c.Assert(show.InAccount, Equals, false)
}

func (s *MySuite) TestShowsVideos(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	videos, err := bs.ShowsVideos(1, 0)
	c.Assert(err, IsNil)
	c.Assert(len(videos), Equals, 6)
	c.Assert(strings.Contains(videos[0].YoutubeURL, "http"), Equals, true)

	videos, err = bs.ShowsVideos(0, 1)
	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, &errAPI{
		[]errorsAPI{err4001},
	})
	c.Assert(len(videos), Equals, 0)

	videos, err = bs.ShowsVideos(1, 1)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoSingleIDUsed)
	c.Assert(len(videos), Equals, 0)

	videos, err = bs.ShowsVideos(0, 0)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errIDNotProperlySet)
	c.Assert(len(videos), Equals, 0)
}
