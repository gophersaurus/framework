Gophersaurus
[![GoDoc](http://godoc.org/github.com/gophersaurus/gf.v1?status.png)](http://godoc.org/github.com/gophersaurus/gf.v1) [![Build Status](https://travis-ci.org/gophersaurus/gf.v1.svg?branch=master)](https://travis-ci.org/gophersaurus/gf.v1) [![Coverage](http://gocover.io/_badge/github.com/gophersaurus/gf.v1?2)](http://gocover.io/github.com/gophersaurus/gf.v1)
============

Gophersaurus is a framework for building API based services quickly. It provides
a structured scaffold for building golang services and abstracts away common
logic away when building API services.

 Gophersaurus was forged in the fires of production, since no API framework
 had yet given us a simple MVC scaffold.

Installation:
-------------
> Note: These instructions assume you already have your $GOPATH configured and
> you are using git with github or github-enterprise.

###Option 1.
Clone or fork this repository locally and start hacking.

###Option 2.
Use the `gf` command line tool to create a new project.

  - `go get github.com/gophersaurus/gf`
  - `gf new project-name`

####Extra Points
Run `go install github.com/mattn/go-sqlite3`. You don't want to build this C
package every time. You have now saved yourself from slow C compile times.

Directory Structure
-------------------
```
.
├── app
│   ├── bootstrap
│   ├── controllers
│   ├── middleware
│   ├── models
│   ├── services
│   └── templates
├── commands
└── public
```

###The Root Directory
The root directory of a fresh Gophersaurus installation contains a variety of
directories.

* The `app` directory contains the core code of your application.
* The `commands` directory contains all the code for CLI commands and flags.
  Gophersaurus makes good use of [`cobra`](https://github.com/spf13/cobra) for
  CLI management.
* The `public` directory contains your static assets such as images, JavaScript
  files, CSS, etc.

####The App Directory

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
* The `templates` directory contains golang templates usually used for rendering
  HTML pages.

Configuration Settings
----------------------
Gophersaurus utilizes the awesome [`viper`](https://github.com/spf13/viper)
package for configuration management.

By default Gophersaurus will search in your projects `app` directory for a file
named `config` to read application settings. [`Viper`](https://github.com/spf13/viper)
supports `JSON`, `YAML`, and `TOML`. You can specify a different file by passing
in the `-c=path/to/file.yml`.

[`Viper`](https://github.com/spf13/viper) also provides support for `etcd` and
`consul`.

An example `app/config.yml` configuration file is provided below.

```YAML
port: 5225
keys:
  gophersaurus:
    - all
databases:
  dbUsers:
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
}L   map[string]DB
}

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
* Send us a pull request against the `develop` branch.

Thanks! :)

License
-------
The Gophersaurus project is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT).
