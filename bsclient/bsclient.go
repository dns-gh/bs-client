// Package bsclient implements a web client for the betaseries API
// - original code from https://github.com/zaibon/tget
package bsclient

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
	baseURL string
	version string
	key     string
	token   *token
}

func (bs *BetaSeries) getToken() (string, error) {
	if bs.token != nil {
		return bs.token.Token, nil
	}
	return "", errNoToken
}

// NewBetaseriesClient creates a betaseries web client
func NewBetaseriesClient(key, login, password string) (*BetaSeries, error) {
	bs := &BetaSeries{
		version: bsVersion,
		baseURL: bsBaseURL,
		key:     key,
	}
	// basic authentication.
	// TODO: OAUTH 2.0
	err := bs.retrieveToken(login, password)
	return bs, err
}

func (bs *BetaSeries) doGet(u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("error creating request for %s: %v", u.String(), err.Error())
		return nil, err
	}

	resp, err := bs.do(req)
	if err != nil {
		log.Printf("error making GET request %s: %v", u.String(), err.Error())
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		apiErr := decodeErr(resp.Body)
		log.Println(apiErr.Error())
		return nil, apiErr
	}
	return resp, nil
}

func (bs *BetaSeries) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-BetaSeries-Version", bs.version)
	req.Header.Set("X-BetaSeries-Key", bs.key)
	if bs.token != nil {
		req.Header.Set("X-BetaSeries-Token", bs.token.Token)
	}

	return http.DefaultClient.Do(req)
}

func (bs *BetaSeries) decode(data interface{}, resp *http.Response, usedAPI, query string) error {
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		log.Printf("Error decoding using '%s' API for '%s' query :%v", usedAPI, query, err)
		return err
	}
	return nil
}

func decodeErr(r io.Reader) *errAPI {
	err := &errAPI{}
	if jsonerr := json.NewDecoder(r).Decode(&err); jsonerr != nil {
		log.Fatalf("Error decoding API error : %v", jsonerr)
	}
	return err
}

func (bs *BetaSeries) retrieveToken(login, password string) error {
	if len(login) == 0 || len(password) == 0 {
		return nil
	}

	u, err := url.Parse(bs.baseURL + "/members/auth")
	if err != nil {
		log.Fatalln(err)
	}
	q := u.Query()
	q.Set("login", login)
	q.Set("password", fmt.Sprintf("%x", md5.Sum([]byte(password))))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return fmt.Errorf("Error creating request for %s: %v\n", u.String(), err.Error())
	}

	resp, err := bs.do(req)
	if err != nil {
		return fmt.Errorf("Error getting token :%v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(decodeErr(resp.Body).Error())
	}

	tokenData := &token{}
	if err := json.NewDecoder(resp.Body).Decode(tokenData); err != nil {
		return fmt.Errorf("Error decoding token response :%v\n", err)
	}
	bs.token = tokenData
	return nil
}
