framework [![GoDoc](http://godoc.org/github.com/gophersaurus/gf.v1?status.png)](http://godoc.org/github.com/gophersaurus/gf.v1) [![Build Status](https://travis-ci.org/gophersaurus/gf.v1.svg?branch=master)](https://travis-ci.org/gophersaurus/gf.v1) [![Go Report Card](http://goreportcard.com/badge/gophersaurus/framework)](http://goreportcard.com/report/gophersaurus/framework)
=========================

Gophersaurus is a MVC framework for building RESTful API services quickly.
It provides a structured scaffolding for abstracting away some common logic.

Gophersaurus was forged for production use.

Installation:
-------------
> Note: These instructions assume you already have your $GOPATH configured.

The easiest way to create a new project is with the `gf` command.

  - `go get -u github.com/gophersaurus/gf`
  - `gf new project-name`

#### Extra Points
  - `go install github.com/mattn/go-sqlite3`.
If your using a SQL database you might not want to build this C package every time.  You have now saved yourself from slow C compile times.

Directory Structure
-------------------

> Running the `$ tree -d` command will show the following result.

```bash
→ tree -d
.
├── app
│   ├── controllers
│   ├── middleware
│   ├── models
│   └── templates
├── cmd
├── public
│   ├── docs
│   │   └── api
│   └── images
└── vendor
```

### The Root Directory
The root directory of a fresh installation contains a variety of directories.

* The `app` directory contains the core code of your application.
* The `cmd` directory contains all the code for CLI commands and flags.
  Gophersaurus makes good use of [`cobra`](https://github.com/spf13/cobra) for
  CLI management.
* The `public` directory contains your static assets such as images, JavaScript
  files, CSS, etc.

#### The App Directory

The "meat" of your application lives in the `app` directory. The `app` directory
ships with a variety of additional directories such as `bootstrap`,
`controllers`, `middleware`, `models`, `services` and `templates`.

* The `bootstrap` directory contains a few files that bootstrap configuration
  settings, databases, and documentation.  
* The `controllers` directory contains all the core business logic of your
  application.  
* The `middleware` directory contains middleware methods that implement the
  `ServeHTTP` interface.  
* The `models` directory contains data object that used usually for marshaling
  and unmarshaling data.
* The `services` directory contains packages that interface with other external
  services. This keeps business logic separate from other service integrations.
  This directory is optional and not created for you by default.
* The `templates` directory contains golang templates usually used for rendering
  HTML pages.

Configuration Settings
----------------------
Gophersaurus utilizes the awesome [`viper`](https://github.com/spf13/viper)
package for configuration management.

By default Gophersaurus will search in your projects root directory for a file
named `.env.yml` to read application settings. [`Viper`](https://github.com/spf13/viper)
supports `JSON`, `YAML`, and `TOML`.

[`Viper`](https://github.com/spf13/viper) also provides support for `etcd` and
`consul`.

An example `.env.yml` configuration file is provided below.

```YAML
port: 5225
keys:
  gophersaurus:
    - all
databases:
  type: mongo
  user: gf
  pass: g0phersaurus!
  address: 10.44.222.233:27017
```

> Note: The port for Gophersaurus applications defaults to `5225`.  
>
> After specifying a key you must specify a whitelist of allowed referer
> addresses. Referer `localhost` values translate as `::1`.
>
> Gophersaurus attempts to compensate for proxies or load balancers by searching
the `Header` for a `X-FORWARDED-FOR` value.

The Database Administrator
--------------------------
We never were overly keen on ORMs. We understand ORMs provide convenience, but
in the new world of SQL and NoSQL databases, no singular ORM solves everything.

Thus the `dba` package (Database Administrator) was born to provide
functionality for both SQL and NoSQL databases.

Here is what the `DatabaseAdmin` struct and `Database` interface looks like:
```Go
type DatabaseAdmin struct {
    Mongo []Database
    SQL   []Database
}
```

```Go
type Database interface {
    Dial(name string) error
    Name() string
    Close()
}
```

The `DatabaseAdmin` object is still something like an ORM. Sorta/Kinda/Maybe.
Right now it implements the `mgo` package for MongoDB and the `gorp` package for
SQL databases. `DatabaseAdmin` keeps both separate, but both implement the same
`Database` interface for convenience.  

Currently the Gophersaurus Framework has support the databases below.
* MongoDB
* PostgreSQL
* MySQL
* SQLite3

Style Guide
-----------
The Gophersaurus project attempts to produce idiomatic go code by utilizing the following standards and style guides.

###Style Guides
- [Golang Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [What's in a name?](http://talks.golang.org/2014/names.slide#1)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Readability](https://talks.golang.org/2014/readability.slide#1)

###Tools
- [`go fmt`](https://golang.org/cmd/gofmt/)
- [`go vet`](https://godoc.org/golang.org/x/tools/cmd/vet)
- [`golint`](https://github.com/golang/lint)
- [`grind`](http://godoc.org/rsc.io/grind)

Contribution guidelines
-----------------------
* Submit an issue.  
* Send us a pull request.

Thanks! :)

License
-------
The Gophersaurus project is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT).
