package bsclient

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var (
	errNoShowsFound      = errors.New("no shows found")
	errNoCharactersFound = errors.New("no characters found")
	errNoVideosFound     = errors.New("no videos found")
	errNoSingleIDUsed    = errors.New("no single id used")
	errIDNotProperlySet  = errors.New("id not properly set")
)

type seasonDetails struct {
	Number   int `json:"number"`
	Episodes int `json:"episodes"`
}

// Show represents the show data returned by the betaserie API
type Show struct {
	// used in episodes/... and shows/... API endpoints
	ID        int    `json:"id"`
	ThetvdbID int    `json:"thetvdb_id"`
	ImdbID    string `json:"imdb_id"`
	Title     string `json:"title"`
	// specific to shows/... API endpoints
	Description    string          `json:"description"`
	Seasons        string          `json:"seasons"`
	SeasonsDetails []seasonDetails `json:"seasons_details"`
	Episodes       string          `json:"episodes"`
	Followers      string          `json:"followers"`
	Comments       string          `json:"comments"`
	Similars       string          `json:"similars"`
	Characters     string          `json:"characters"`
	Creation       string          `json:"creation"`
	Genres         []string        `json:"genres"`
	Length         string          `json:"length"`
	Network        string          `json:"network"`
	Rating         string          `json:"rating"`
	Status         string          `json:"status"`
	Language       string          `json:"language"`
	Notes          struct {
		Total int     `json:"total"`
		Mean  float32 `json:"mean"`
		User  int     `json:"user"`
	} `json:"notes"`
	InAccount bool `json:"in_account"`
	Images    struct {
		Show   string `json:"show"`
		Banner string `json:"banner"`
		Box    string `json:"box"`
		Poster string `json:"poster"`
	} `json:"images"`
	Aliases []string `json:"aliases"`
	User    struct {
		Archived  bool `json:"archived"`
		Favorited bool `json:"favorited"`
		// remove temporarily until resolved by betaseries API: https://www.betaseries.com/bugs/api/383
		/*Remaining int    `json:"remaining"`
		Status    int    `json:"status"`
		Last      string `json:"last"`
		Tags      string `json:"tags"`
		*/
	} `json:"user"`
	ResourceURL string `json:"resource_url"`
	// specific to episodes/... API endpoints
	Remaining int       `json:"remaining"`
	Unseen    []Episode `json:"unseen"`
}

type shows struct {
	Shows  []Show        `json:"shows"`
	Errors []interface{} `json:"errors"`
}

type showItem struct {
	Show   *Show         `json:"show"`
	Errors []interface{} `json:"errors"`
}

func (bs *BetaSeries) doGetShows(u *url.URL, usedAPI string) ([]Show, error) {
	resp, err := bs.do("GET", u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &shows{}
	err = bs.decode(data, resp, usedAPI, u.RawQuery)
	if err != nil {
		return nil, err
	}

	if len(data.Shows) < 1 {
		return nil, errNoShowsFound
	}

	return data.Shows, nil
}

// ShowsSearch returns a slice of shows found with the given query
// The slice is of size 100 maximum and the results are ordered by popularity by default.
func (bs *BetaSeries) ShowsSearch(query string) ([]Show, error) {
	usedAPI := "/shows/search"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	q.Set("title", strings.ToLower(query))
	q.Set("order", "popularity")
	q.Set("nbpp", "100")
	u.RawQuery = q.Encode()

	return bs.doGetShows(u, usedAPI)
}

// ShowsRandom returns a slice of random shows. The maximum size of the slice is given
// by the 'num' parameter. If you want to get only summarized info, use the 'summary parameter.
func (bs *BetaSeries) ShowsRandom(num int, summary bool) ([]Show, error) {
	usedAPI := "/shows/random"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	if num >= 0 {
		q.Set("nb", strconv.Itoa(num))
	}
	if summary {
		q.Set("summary", strconv.FormatBool(summary))
	}
	u.RawQuery = q.Encode()

	return bs.doGetShows(u, usedAPI)
}

// Character represents the character data returned by the betaserie API
type Character struct {
	ID          int    `json:"id"`
	ShowID      int    `json:"show_id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Actor       string `json:"actor"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

type characters struct {
	Characters []Character `json:"characters"`
}

// ShowsCharacters returns a slice of shows found with the given ID
// The slice is of size 100 maximum and the results are ordered by popularity by default.
func (bs *BetaSeries) ShowsCharacters(id int) ([]Character, error) {
	usedAPI := "/shows/characters"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	q.Set("id", strconv.Itoa(id))
	u.RawQuery = q.Encode()

	resp, err := bs.do("GET", u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &characters{}
	err = bs.decode(data, resp, usedAPI, u.RawQuery)
	if err != nil {
		return nil, err
	}

	if len(data.Characters) < 1 {
		return nil, errNoCharactersFound
	}

	return data.Characters, nil
}

// ShowsList returns a slice of shows from an interval. It can return every shows if wanted.
// 'since' : only displays shows from a specified data (timestamp UNIX - optional)
// 'starting' : only displays shows beginning with the specified string (optional)
// 'start' : show id number to begin the listing with (default 0, optional)
// 'limit' : maximum size of the returned slice (default to everything, optional)
func (bs *BetaSeries) ShowsList(since, starting string, start, limit int) ([]Show, error) {
	usedAPI := "/shows/list"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	q.Set("order", "popularity")
	if len(since) != 0 {
		q.Set("since", since)
	}
	if len(starting) != 0 {
		q.Set("starting", starting)
	}
	if start > 0 {
		q.Set("start", strconv.Itoa(start))
	}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	u.RawQuery = q.Encode()

	return bs.doGetShows(u, usedAPI)
}

func (bs *BetaSeries) showUpdate(method, endoint string, id int) (*Show, error) {
	usedAPI := "/shows/" + endoint
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	q.Set("id", strconv.Itoa(id))
	u.RawQuery = q.Encode()

	resp, err := bs.do(method, u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	show := &showItem{}
	err = bs.decode(show, resp, usedAPI, u.RawQuery)
	if err != nil {
		return nil, err
	}

	return show.Show, nil
}

// ShowAdd adds the show represented by the given 'id' to the user's account.
func (bs *BetaSeries) ShowAdd(id int) (*Show, error) {
	return bs.showUpdate("POST", "show", id)
}

// ShowRemove removes the show represented by the given 'id' from user's account.
func (bs *BetaSeries) ShowRemove(id int) (*Show, error) {
	return bs.showUpdate("DELETE", "show", id)
}

// ShowArchive archives the show represented by the given 'id' from user's account
func (bs *BetaSeries) ShowArchive(id int) (*Show, error) {
	return bs.showUpdate("POST", "archive", id)
}

// ShowNotArchive remove from archives the show represented by the given 'id' from user's account
func (bs *BetaSeries) ShowNotArchive(id int) (*Show, error) {
	return bs.showUpdate("DELETE", "archive", id)
}

// Video represents the video data returned by the betaserie API
type Video struct {
	ID         int    `json:"id"`
	ShowID     int    `json:"show_id"`
	YoutubeID  string `json:"youtube_id"`
	YoutubeURL string `json:"youtube_url"`
	Title      string `json:"title"`
	Season     int    `json:"season"`
	Episode    int    `json:"episode"`
	Login      string `json:"login"`
	LoginID    int    `json:"login_id"`
}

type videos struct {
	Videos []Video       `json:"videos"`
	Errors []interface{} `json:"errors"`
}

// ShowsVideos returns a slice of videos added by the betaseries members
// on a specific show using the show 'id' or 'tvdbID' (strictly positive)
// Note: do not use both ids, it will return an error
func (bs *BetaSeries) ShowsVideos(id, tvdbID int) ([]Video, error) {
	usedAPI := "/shows/videos"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	if id > 0 && tvdbID > 0 {
		return nil, errNoSingleIDUsed
	} else if id > 0 {
		q.Set("id", strconv.Itoa(id))
	} else if tvdbID > 0 {
		q.Set("thetvdb_id ", strconv.Itoa(tvdbID))
	} else {
		return nil, errIDNotProperlySet
	}
	u.RawQuery = q.Encode()

	resp, err := bs.do("GET", u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &videos{}
	err = bs.decode(data, resp, usedAPI, u.RawQuery)
	if err != nil {
		return nil, err
	}

	if len(data.Videos) < 1 {
		return nil, errNoVideosFound
	}

	return data.Videos, nil
}
