package bsclient

import (
	"os"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestPicturesShows(c *C) {
	key := os.Getenv("BS_API_KEY")
	bs, err := NewBetaseriesClient(key, "Dev050", "developer")
	c.Assert(err, IsNil)
	picture, err := bs.PicturesShows(1, -1, -1)
	c.Assert(err, IsNil)
	c.Assert(len(picture), Equals, 50241)

	picture, err = bs.PicturesShows(1, 100, 100)
	c.Assert(err, IsNil)
	c.Assert(len(picture), Equals, 10985)

	picture, err = bs.PicturesShows(0, 100, 100)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errIDMustBeStrictlyPositive)
}
