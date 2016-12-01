package bsclient

import (
	"errors"
	"net/url"
	"strconv"
)

var (
	errNoNewsFound = errors.New("no news found")
)

// News represents a news of a particular tv show
type News struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	PictureURL string `json:"picture_url"`
	Date       string `json:"date"`
}

type news struct {
	News   []News        `json:"news"`
	Errors []interface{} `json:"errors"`
}

// NewsLast returns a slice of news of tv shows
// If 'number' is strictly negative, it returns a default of 10 news maximum.
// THe 'tailored' parameter returns tv show news of the identified member.
func (bs *BetaSeries) NewsLast(number int, tailored bool) ([]News, error) {
	usedAPI := "/planning/general"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return nil, errURLParsing
	}
	q := u.Query()
	q.Set("number", strconv.Itoa(number))
	q.Set("tailored", strconv.FormatBool(tailored))
	u.RawQuery = q.Encode()

	resp, err := bs.doGet(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &news{}
	err = bs.decode(data, resp, usedAPI, u.RawQuery)
	if err != nil {
		return nil, err
	}

	if len(data.News) < 1 {
		return nil, errNoNewsFound
	}

	return data.News, nil
}
