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
	errURLParsing        = errors.New("url parsing error")
)

type seasonDetails struct {
	Number   int `json:"number"`
	Episodes int `json:"episodes"`
}

// Show represents the show data returned by the betaserie API
type Show struct {
	ID             int             `json:"id"`
	ThetvdbID      int             `json:"thetvdb_id"`
	ImdbID         string          `json:"imdb_id"`
	Title          string          `json:"title"`
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
		Total string `json:"total"`
		Mean  string `json:"mean"`
		User  int    `json:"user"`
	} `json:"notes"`
	InAccount bool `json:"in_account"`
	Images    struct {
		Poster string `json:"poster"`
	} `json:"images"`
	Aliases []interface{} `json:"aliases"`
	User    struct {
		Archived  bool        `json:"archived"`
		Favorited bool        `json:"favorited"`
		Remaining int         `json:"remaining"`
		Status    int         `json:"status"`
		Last      string      `json:"last"`
		Tags      interface{} `json:"tags"`
	} `json:"user"`
	ResourceURL string `json:"resource_url"`
}

type shows struct {
	Shows []Show `json:"shows"`
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

	resp, err := bs.doGet(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &shows{}
	err = bs.decode(data, resp, usedAPI, query)
	if err != nil {
		return nil, err
	}

	if len(data.Shows) < 1 {
		return nil, errNoShowsFound
	}

	return data.Shows, nil
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
	query := strconv.Itoa(id)
	q.Set("id", query)
	u.RawQuery = q.Encode()

	resp, err := bs.doGet(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &characters{}
	err = bs.decode(data, resp, usedAPI, query)
	if err != nil {
		return nil, err
	}

	if len(data.Characters) < 1 {
		return nil, errNoCharactersFound
	}

	return data.Characters, nil
}
