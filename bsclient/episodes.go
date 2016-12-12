package bsclient

import (
	"net/url"
	"strconv"
)

// EpisodesList returns a slice of unseen episodes ordered by shows.
func (bs *BetaSeries) EpisodesList(showID, userID int) ([]Show, error) {
	usedAPI := "/episodes/list"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	if showID >= 0 {
		q.Set("showId", strconv.Itoa(showID))
	}
	if userID >= 0 {
		q.Set("userId", strconv.Itoa(userID))
	}
	u.RawQuery = q.Encode()

	return bs.doGetShows(u, usedAPI)
}
