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

func (s *MySuite) TestEpisodesList(c *C) {
	bs, key, id := makeClientAndAddShow(c)
	shows, err := bs.EpisodesList(id, -1)
	c.Assert(err, IsNil)
	c.Assert(shows, HasLen, 1)

	show, err := bs.ShowRemove(id)
	c.Assert(err, IsNil)
	c.Assert(show.InAccount, Equals, false)

	_, err = bs.EpisodesList(-1, -1)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoShowsFound)

	bs, err = NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	_, err = bs.EpisodesList(0, 0)
	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, &errAPI{
		[]errorsAPI{err2001},
	})
}

func (s *MySuite) TestEpisodesDownloaded(c *C) {
	bs, _, id := makeClientAndAddShow(c)
	shows, err := bs.EpisodesList(id, -1)
	c.Assert(err, IsNil)
	c.Assert(shows, HasLen, 1)
	c.Assert(shows[0].Unseen, HasLen, 62)

	episode, err := bs.EpisodeDownloaded(shows[0].Unseen[0].ID)
	c.Assert(err, IsNil)
	c.Assert(episode.User.Downloaded, Equals, true)

	episode, err = bs.EpisodeNotDownloaded(shows[0].Unseen[0].ID)
	c.Assert(err, IsNil)
	c.Assert(episode.User.Downloaded, Equals, false)

	show, err := bs.ShowRemove(id)
	c.Assert(err, IsNil)
	c.Assert(show.InAccount, Equals, false)
}
