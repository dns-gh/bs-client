package bsclient

import (
	"bytes"
	"errors"
	"io"
	"net/url"
	"strconv"
)

var (
	errIDMustBeStrictlyPositive = errors.New("id must be strictly positive")
)

// PicturesShows returns a picture of the tv show identified by 'id'.
// If 'id' is negative, a default betaseries picture will be returned with an error code.
// The optional 'width' and 'height' parameters must be both strictly
// positive in order to be used.
func (bs *BetaSeries) PicturesShows(id, width, height int) (string, error) {
	usedAPI := "/pictures/shows"
	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		return "", errURLParsing
	}
	q := u.Query()
	if id <= 0 {
		return "", errIDMustBeStrictlyPositive
	}
	q.Set("id", strconv.Itoa(id))
	if width > 0 && height > 0 {
		q.Set("width", strconv.Itoa(width))
		q.Set("height", strconv.Itoa(height))
	}
	u.RawQuery = q.Encode()
	resp, err := bs.do("GET", u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return "", err
	}

	return buf.String(), nil
}
