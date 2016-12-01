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
$ go test ...bsclient -gocheck.vv -test.v -gocheck.f Test
=== RUN   Test
START: bsclient_test.go:17: MySuite.TestNewBS
PASS: bsclient_test.go:17: MySuite.TestNewBS    0.001s

START: bsclient_test.go:48: MySuite.TestNewBSGetToken
PASS: bsclient_test.go:48: MySuite.TestNewBSGetToken    0.001s

START: bsclient_test.go:38: MySuite.TestNewBSGetTokenWithAPIKey
PASS: bsclient_test.go:38: MySuite.TestNewBSGetTokenWithAPIKey  0.116s

START: bsclient_test.go:27: MySuite.TestNewBSGetTokenWithoutAPIKey
PASS: bsclient_test.go:27: MySuite.TestNewBSGetTokenWithoutAPIKey       0.040s

START: news_test.go:10: MySuite.TestNewsLast
PASS: news_test.go:10: MySuite.TestNewsLast     1.965s

START: planning_test.go:24: MySuite.TestPlanningGeneral
PASS: planning_test.go:24: MySuite.TestPlanningGeneral  0.283s

START: planning_test.go:40: MySuite.TestPlanningIncoming
PASS: planning_test.go:40: MySuite.TestPlanningIncoming 0.091s

START: planning_test.go:53: MySuite.TestPlanningMember
PASS: planning_test.go:53: MySuite.TestPlanningMember   0.144s

START: planning_test.go:90: MySuite.TestPlanningMemberWithCredentials
PASS: planning_test.go:90: MySuite.TestPlanningMemberWithCredentials    0.258s

START: shows_test.go:37: MySuite.TestShowsCharacters
PASS: shows_test.go:37: MySuite.TestShowsCharacters     0.178s

START: shows_test.go:20: MySuite.TestShowsSearch
PASS: shows_test.go:20: MySuite.TestShowsSearch 0.096s

OK: 11 passed
--- PASS: Test (3.19s)
PASS
ok      github.com/dns-gh/bs-client/bsclient    3.271s
```

## LICENSE

See included LICENSE file.