# bs-client

[![Go Report Card](https://goreportcard.com/badge/github.com/dns-gh/bs-client)](https://goreportcard.com/report/github.com/dns-gh/bs-client)

[![GoDoc](https://godoc.org/github.com/dns-gh/bs-client?status.png)]
(https://godoc.org/github.com/dns-gh/bs-client)

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

TODO

## LICENSE

See included LICENSE file.