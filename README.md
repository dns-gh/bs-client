# bs-client

[![Go Report Card](https://goreportcard.com/badge/github.com/dns-gh/bs-client)](https://goreportcard.com/report/github.com/dns-gh/bs-client)

[![GoDoc](https://godoc.org/github.com/dns-gh/bs-client/bsclient?status.png)]
(https://godoc.org/github.com/dns-gh/bs-client/bsclient)

bs-client is a web client for the betaseries website API

## Motivation

Used in the Betaseries Twitter Bot project: https://github.com/dns-gh/bsbot

## Installation

- It requires Go language of course. You can set it up by downloading it here: https://golang.org/dl/
- Install it here C:/Go.
- Set your GOPATH, GOROOT and PATH environment variables:

```
export GOROOT=C:/Go
export GOPATH=WORKING_DIR
export PATH=C:/Go/bin:${PATH}
```

- Download and install the package:

```
@working_dir $ go get github.com/dns-gh/bs-client/...
@working_dir $ go install github.com/dns-gh/bs-client/bsclient
```

## Example

See the https://github.com/dns-gh/bsbot

## Tests

Example of a test launch:
```
$ export BS_API_KEY=YOUR_BETASERIES_KEY && go test ...bsclient -gocheck.vv -test.v -gocheck.f Test
=== RUN   Test
START: episodes_test.go:37: MySuite.TestEpisodesDownloaded
PASS: episodes_test.go:37: MySuite.TestEpisodesDownloaded       0.951s

START: episodes_test.go:14: MySuite.TestEpisodesList
PASS: episodes_test.go:14: MySuite.TestEpisodesList     0.731s

START: episodes_test.go:57: MySuite.TestEpisodesWatched
PASS: episodes_test.go:57: MySuite.TestEpisodesWatched  0.752s

START: bsclient_test.go:17: MySuite.TestNewBS
PASS: bsclient_test.go:17: MySuite.TestNewBS    0.000s

START: bsclient_test.go:50: MySuite.TestNewBSGetToken
PASS: bsclient_test.go:50: MySuite.TestNewBSGetToken    0.001s

START: bsclient_test.go:40: MySuite.TestNewBSGetTokenWithAPIKey
PASS: bsclient_test.go:40: MySuite.TestNewBSGetTokenWithAPIKey  0.146s

START: bsclient_test.go:28: MySuite.TestNewBSGetTokenWithoutAPIKey
PASS: bsclient_test.go:28: MySuite.TestNewBSGetTokenWithoutAPIKey       0.077s

START: news_test.go:10: MySuite.TestNewsLast
PASS: news_test.go:10: MySuite.TestNewsLast     0.095s

START: pictures_test.go:9: MySuite.TestPicturesShows
PASS: pictures_test.go:9: MySuite.TestPicturesShows     0.271s

START: planning_test.go:24: MySuite.TestPlanningGeneral
PASS: planning_test.go:24: MySuite.TestPlanningGeneral  0.433s

START: planning_test.go:40: MySuite.TestPlanningIncoming
PASS: planning_test.go:40: MySuite.TestPlanningIncoming 0.140s

START: planning_test.go:53: MySuite.TestPlanningMember
PASS: planning_test.go:53: MySuite.TestPlanningMember   0.211s

START: planning_test.go:90: MySuite.TestPlanningMemberWithCredentials
PASS: planning_test.go:90: MySuite.TestPlanningMemberWithCredentials    0.359s

START: shows_test.go:57: MySuite.TestShowsCharacters
PASS: shows_test.go:57: MySuite.TestShowsCharacters     0.193s

START: shows_test.go:75: MySuite.TestShowsList
PASS: shows_test.go:75: MySuite.TestShowsList   0.286s

START: shows_test.go:38: MySuite.TestShowsRandom
PASS: shows_test.go:38: MySuite.TestShowsRandom 0.209s

START: shows_test.go:21: MySuite.TestShowsSearch
PASS: shows_test.go:21: MySuite.TestShowsSearch 0.202s

START: shows_test.go:109: MySuite.TestShowsUpdate
PASS: shows_test.go:109: MySuite.TestShowsUpdate        0.604s

START: shows_test.go:135: MySuite.TestShowsVideos
PASS: shows_test.go:135: MySuite.TestShowsVideos        0.131s

OK: 19 passed
--- PASS: Test (5.85s)
PASS
ok      github.com/dns-gh/bs-client/bsclient    6.018s
```

## LICENSE

See included LICENSE file.