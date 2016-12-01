package bsclient

import (
	"os"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestNewsLast(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "", "")
	c.Assert(err, IsNil)
	news, err := bs.NewsLast(1, false)
	if len(news) > 0 {
		c.Assert(err, IsNil)
		c.Assert(len(news), Equals, 1)
		c.Assert(len(news[0].Date), Equals, 19)
		c.Assert(strings.Contains(news[0].URL, "http"), Equals, true)
		c.Assert(strings.Contains(news[0].PictureURL, "http"), Equals, true)
	} else {
		c.Assert(err, NotNil)
		c.Assert(err, Equals, errNoNewsFound)
	}
}
