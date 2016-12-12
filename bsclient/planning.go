package bsclient

import (
	"errors"
	"net/url"
	"strconv"
)

var (
	errNoEpisodesFound = errors.New("no episodes found")
)

// Episode represents the episode data returned by the betaserie API
type Episode struct {
	ID        int    `json:"id"`
	ThetvdbID int    `json:"thetvdb_id"`
	YoutubeID string `json:"youtube_id"`
	Title     string `json:"title"`
	Season    int    `json:"season"`
	Episode   int    `json:"episode"`
	Show      struct {
		ID        int    `json:"id"`
		ThetvdbID int    `json:"thetvdb_id"`
		Title     string `json:"title"`
	} `json:"show"`
	Code        string `json:"code"`
	Global      int    `json:"global"`
	Special     int    `json:"special"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Note        struct {
		Total string `json:"total"`
		Mean  string `json:"mean"`
		User  int    `json:"user"`
	} `json:"note"`
	User struct {
		Seen       bool `json:"seen"`
		Downloaded bool `json:"downloaded"`
	} `json:"user"`
	Comments  string        `json:"comments"`
	Subtitles []interface{} `json:"subtitles"`
}

type episodes struct {
	Episodes []Episode     `json:"episodes"`
	Errors   []interface{} `json:"errors"`
}

func (bs *BetaSeries) doGetPlanning(u *url.URL, usedAPI string) ([]Episode, error) {
	resp, err := bs.do("GET", u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &episodes{}
	err = bs.decode(data, resp, usedAPI, u.RawQuery)
	if err != nil {
		return nil, err
	}

	if len(data.Episodes) < 1 {
		return nil, errNoEpisodesFound
	}

	return data.Episodes, nil
}

// PlanningGeneral returns a slice of episodes found in [date-before, date+after] timeline.
// Note: the 'date' input must be in YYYY-MM-JJ format
func (bs *BetaSeries) PlanningGeneral(date string, before, after int) ([]Episode, error) {
	usedAPI := "/planning/general"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	q.Set("date", date)
	q.Set("before", strconv.Itoa(before))
	q.Set("after", strconv.Itoa(after))
	u.RawQuery = q.Encode()
	return bs.doGetPlanning(u, usedAPI)
}

// PlanningIncoming returns a slice of the first episodes of each tv show
// that are about to be broacasted.
func (bs *BetaSeries) PlanningIncoming() ([]Episode, error) {
	usedAPI := "/planning/incoming"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	return bs.doGetPlanning(u, usedAPI)
}

// PlanningMember returns a slice of episodes of the member 'id'
// If 'id' is not set to a proper member it returns an error.
// If 'id' is strictly negative, the returned episodes
// are the ones of the identified member doing the request.
// The parameter 'unseen' filters not seen episodes.
// The parameter 'month' filters episodes of the given month with the format YYYY-MM.
// Note: the 'month' value can be the string "now".
func (bs *BetaSeries) PlanningMember(id int, unseen bool, month string) ([]Episode, error) {
	usedAPI := "/planning/member"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	if id >= 0 {
		q.Set("id", strconv.Itoa(id))
	}
	q.Set("unseen", strconv.FormatBool(unseen))
	q.Set("month", month)
	u.RawQuery = q.Encode()
	return bs.doGetPlanning(u, usedAPI)
}
