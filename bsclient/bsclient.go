// Package bsclient implements a web client for the betaseries API
package bsclient

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	bsBaseURL = "https://api.betaseries.com"
	bsVersion = "2.4"
)

var (
	errNoToken    = errors.New("no token")
	errURLParsing = errors.New("url parsing error")
)

type errorsAPI struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// errAPI represents an error returned by the API
type errAPI struct {
	Errors []errorsAPI `json:"errors"`
}

func (e *errAPI) Error() string {
	out := ""
	for _, e := range e.Errors {
		out += fmt.Sprintf("%s\n", e.Text)
	}
	return out
}

// token is a struct return by the betaseries API when requesting a token
type token struct {
	User struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		InAccount bool   `json:"in_account"`
	} `json:"user"`
	Token  string        `json:"token"`
	Hash   string        `json:"hash"`
	Errors []interface{} `json:"errors"`
}

// BetaSeries represents the web client to the BetaSeries API
type BetaSeries struct {
	baseURL    string
	version    string
	key        string
	token      *token
	httpClient *http.Client
}

func (bs *BetaSeries) getToken() (string, error) {
	if bs.token != nil {
		return bs.token.Token, nil
	}
	return "", errNoToken
}

// NewBetaseriesClient creates a betaseries web client
func NewBetaseriesClient(key, login, password string) (*BetaSeries, error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	bs := &BetaSeries{
		version: bsVersion,
		baseURL: bsBaseURL,
		key:     key,
		httpClient: &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		},
	}
	// basic authentication.
	// TODO: OAUTH 2.0
	err := bs.retrieveToken(login, password)
	return bs, err
}

func (bs *BetaSeries) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-BetaSeries-Version", bs.version)
	req.Header.Set("X-BetaSeries-Key", bs.key)
	if bs.token != nil {
		req.Header.Set("X-BetaSeries-Token", bs.token.Token)
	}

	return bs.httpClient.Do(req)
}

func (bs *BetaSeries) do(method string, u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := bs.doRequest(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		apiErr := decodeErr(resp.Body)
		return nil, apiErr
	}
	return resp, nil
}

func (bs *BetaSeries) decode(data interface{}, resp *http.Response, usedAPI, query string) error {
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}
	return nil
}

func decodeErr(r io.Reader) *errAPI {
	err := &errAPI{}
	// note that 404 error not found on 'picture, err = bs.PicturesShows(0, 100, 100)' is not handled by errAPI
	json.NewDecoder(r).Decode(&err)
	return err
}

func (bs *BetaSeries) retrieveToken(login, password string) error {
	usedAPI := "/members/auth"
	if len(login) == 0 || len(password) == 0 {
		return nil
	}

	u, err := url.Parse(bs.baseURL + usedAPI)
	if err != nil {
		log.Fatalln(err)
	}
	q := u.Query()
	q.Set("login", login)
	q.Set("password", fmt.Sprintf("%x", md5.Sum([]byte(password))))
	u.RawQuery = q.Encode()

	resp, err := bs.do("POST", u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tokenData := &token{}
	err = bs.decode(tokenData, resp, usedAPI, u.RawQuery)
	if err != nil {
		return err
	}
	bs.token = tokenData
	return nil
}
