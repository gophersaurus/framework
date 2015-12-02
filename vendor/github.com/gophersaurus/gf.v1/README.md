gf.v1 [![GoDoc](http://godoc.org/github.com/gophersaurus/gf.v1?status.png)](http://godoc.org/github.com/gophersaurus/gf.v1) [![Build Status](https://travis-ci.org/gophersaurus/gf.v1.svg?branch=master)](https://travis-ci.org/gophersaurus/gf.v1) [![Go Report Card](http://goreportcard.com/badge/gophersaurus/gf.v1)](http://goreportcard.com/report/gophersaurus/gf.v1)
======

`gf.v1` stands for the gophersaurus framework Version One. Separating out package versions into different repositories helps avoid breaking changes.

Packages
--------

| Package  | Description                   | Coverage |
|:---------|:------------------------------|---------:|
| [config](http://godoc.org/github.com/gophersaurus/gf.v1/config) | configuration settings | [![Coverage](http://gocover.io/_badge/github.com/spf13/viper?1)](http://gocover.io/github.com/spf13/viper) |
| [dba](http://godoc.org/github.com/gophersaurus/gf.v1/dba)      | database connections   |     mongodb tested |
| [docs](http://godoc.org/github.com/gophersaurus/gf.v1/docs)     | generate documentation   | [![Coverage](http://gocover.io/_badge/github.com/gophersaurus/gf.v1/docs?2)](http://gocover.io/github.com/gophersaurus/gf.v1/docs) |
| [http](http://godoc.org/github.com/gophersaurus/gf.v1/http)     | http requests          | [![Coverage](http://gocover.io/_badge/github.com/gophersaurus/gf.v1/http?1)](http://gocover.io/github.com/gophersaurus/gf.v1/http) |
|	[mock](http://godoc.org/github.com/gophersaurus/gf.v1/mock)     | mock http requests            | [![Coverage](http://gocover.io/_badge/github.com/gophersaurus/gf.v1/mock?2.5)](http://gocover.io/github.com/gophersaurus/gf.v1/mock) |
| [render](http://godoc.org/github.com/gophersaurus/gf.v1/render)   | http response formats         | [![Coverage](http://gocover.io/_badge/github.com/gophersaurus/gf.v1/render?1)](http://gocover.io/github.com/gophersaurus/gf.v1/render) |
| [resource](http://godoc.org/github.com/gophersaurus/gf.v1/resource) | CRUD endpoints                 | [![Coverage](http://gocover.io/_badge/github.com/gophersaurus/gf.v1/resource?2)](http://gocover.io/github.com/gophersaurus/gf.v1/resource) |
|	[router](http://godoc.org/github.com/gophersaurus/gf.v1/router)   | multiplex router              | [![Coverage](http://gocover.io/_badge/github.com/gophersaurus/gf.v1/router?2)](http://gocover.io/github.com/gophersaurus/gf.v1/router) |

Install
-------

`go get github.com/gophersaurus/gf.v1`

Update
------

`go get -u github.com/gophersaurus/gf.v1`

Style Guide
-----------

The gophersaurus framework project attempts to produce idiomatic go code by utilizing the following standards and style guides.

### Style Guides
- [Golang Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [What's in a name?](http://talks.golang.org/2014/names.slide#1)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Readability](https://talks.golang.org/2014/readability.slide#1)

### Tools
- [`go fmt`](https://golang.org/cmd/gofmt/)
- [`go vet`](https://godoc.org/golang.org/x/tools/cmd/vet)
- [`golint`](https://github.com/golang/lint)
- [`grind`](http://godoc.org/rsc.io/grind)

License
-------

The gophersaurus framework project is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT).
