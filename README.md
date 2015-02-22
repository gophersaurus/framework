Gophersaurus
=============================

Gophersaurus is a framework for building monolithic web services quickly. It provides a more structured approach to building go services and abstracts much of the common logic away for building an API.  These traits make Gophersaurus a great choice for projects that need to be rapidly developed, yet maintain consistency.  

![gophersaurus](https://git.target.com/gophersaurus/art/raw/master/gophersaurus.png)

Gophersaurus provides structure for large teams of gophers, and wraps many open source golang packages to abstract away common logic.  It is for these reasons Gophersaurus got its name.  Its a big package.

Gophersaurus has been tested in production, it can scale, and basically gets the job done, but it is not the framework for trying to eek out every bit of performance possible.  As the framework matures, it will become more performant, but honestly... if you want the best golang performance possible, don't use a framework.  The go standard library will do.  

Gophersaurus is heavily inspired by other backend frameworks, but especially the Laravel PHP Framework.  Just like Laravel, Gophersaurus encourages the use of Models, Controllers, Resources, and JSON views/responses.

> IMPORTANT NOTE: Gophersaurus is still in development and currently the API is not stable.  We will lockdown the API soon, and then gf.v1 will cease to have breaking changes.  All breaking changes will be diverted gf.v2.

###Installation:

> Note: These instructions assume you already have your $GOPATH configured and you are using git with github or github-enterprise.

####Code Setup

1. Clone this repositiory. `git clone git@git.target.com:gophersaurus/gophersaurus.git`
2. Enter the root project directory. `cd gophersaurus`
4. Run the command `go get ./...` to ensure you have all nessesary dependencies locally.  
   If you want the latest dependencies and verbose mode on run `go get -u -v ./...`.
5. Rename `gophersaurus` to your project name.                          
   First move outside of the project root. `cd ..`                            
   Then change the project directory name. `mv gophersaurus project-name`
6. Do a search and replace in your project root.                                                      
   Move into your project root directory. `cd project-name`                                                 
   Search `git.target.com/gophersaurus/gophersaurus` replace with `your.git.com/your-org/project-name`


####Git Setup

1. In the project directory run `git remote set-url origin git@your.git.com:your-org/your-project-name.git`

####Extra Points

1. Run `go install github.com/mattn/go-sqlite3`.  You don't want to build this C package every time.  
   Save yourself now from slow C compile times.

Now you should be able to run `go build` from your project root and run all manner of `git` commands.

###Directory Structure

```
├── app
│   ├── controllers
│   ├── middleware
│   ├── models
│   └── services
│       └── api.openweathermap.org
│           └── data
│               └── 2.5
├── bootstrap
├── config
└── public
```

####The Root Directory

The root directory of a fresh Laravel installation contains a variety of folders:

* The `app` directory, as you might expect, contains the core code of your application.  
  It also implements the `server` `package` in your application.
* The `bootstrap` folder contains a few files that bootstrap the framework and configure autoloading.  
  It also implements the `bootstrap` `package` in your application.
* The `config` directory, as the name implies, contains all of your application's configuration settings and logic.
  It also implements the `config` `package` in your application.
* The `public` directory contains the front controller and your assets (images, JavaScript, CSS, etc.).

####The App Directory

The "meat" of your application lives in the `app` directory. The `app` directory ships with a variety of additional directories such as `controllers`, `middleware`, `models`, and `services`.

* The `controllers` directory, contains all the core controller code of your application.  
  It also implements the `controllers` `package` in your application.
* The `middleware` directory, contains all the core middleware code of your application.  
  It also implements the `middleware` `package` in your application.
* The `models` directory, contains all the core middleware code of your application.  
  It also implements the `models` `package` in your application.
* The `services` directory, contains all the service code of your application.  
  Multiple service `packages` are implemented in the `services` directory.  Service `package` names usually depend on the kind of service, as well as the URI endpoint for that particular service.  
  
> Note on `service` `package` names: In the example directory structure above, we can deterime that the endpoint for the `weather` `service` is located at `http://api.openweathermap.org/data/2.5/weather`.  This convention is useful for quickly identifying a URI `service` endpoint in your application.

###Configuration Settings

By default Gophersaurus will look in your projects root folder for a `.config.yml` file to read all application settings.  You can specify a different file by passing in the `-c=path/to/your/file.yml`.  Gophersaurus can also read `.json` files instead of `.yml` if you prefer.  

An example `.config.yml` configuration file is provided below:

```YAML
config:
  port: 8080
  keys:
    x78348djas-acceptOnlyTheseRefererKey:
    - 10.87.87.64
    - 34.87.65.10
    x78348djas-acceptOnlyLocalhostKey:
    - localhost
    x78348djas-acceptAnythingKey:
  databases:
  - type: mongo
  name: mongoDatabaseName
  user: mongoUserName
  pass: mongoUserPassword
  address: localhost:27017/mongoDatabaseName
```

> Note: Referer `localhost` values currently translate as `::1`.  Most `/etc/hosts` files have `::1` listed last after `127.0.0.1`.  Also do not attempt to compensate for proxies or loadbalancers unless you know what your doing. 
Gophersaurus will search `Header` for a `X-FORWARDED-FOR` value by default.

###The Database Administrator

Gophersaurus was never overly keen on ORMs.  We understand ORMs are a necessary convenience, but in the new world of SQL and NoSQL databases, no singular Golang ORM has emerged.  (If you do know of a golang ORM to rule them all, please let us know.)

Thus the `DBA` or Database Administrator object was born to help with SQL vs NoSQL.  Here is what it looks like:

```Go
type DBA struct {
	NoSQL map[string]DB
	SQL   map[string]DB
}
```

The `DBA` object in Gophersaurus is still something like an ORM.  Sorta/Kinda/Maybe.  Right now it impliments the `mgo` `package` for NoSQL and the `gorp` `package` for SQL.  The `DBA` keeps both seperate, but they impliment the same `DB` interface for easy access.  

Currently Gophersaurus supports following databases:
* MongoDB
* PostgreSQL
* MySQL
* SQLite3

###Style Guide

We believe it is important for a framework to provide a style guide, not just code.  

Instead of reinventing the wheel, we recommend `gofmt` and `goimports` to automatically format go code properly.  Beyond these awesome tools we also recommend gophers to keep close to the internal Golang Code Review standards at https://github.com/golang/go/wiki/CodeReviewComments.

###What Problem Is Gophersaurus Solving?

We believe in golang there is a need for a large framework to be used in enterprise environments. There are many golang frameworks such as Revel, Traffic, Martini, Gorilla, Goweb, and more. These frameworks are great at what they do, but lack folder structure and strong opinions.

Package flexibility is great when you need to write a small service, but when you start to grow a larger robust codebase it starts to become a nightmare. This is especially true when you have more than one developer.

Our solution/plan has been to steal all of the good directory structure other frameworks are famous for (like Laravel who copied Ruby on Rails), while keeping our own golang code as idiomatic as possible. We aren’t doing anything new, rather we are organizing all the good work the community has already achieved.

###Under The Hood

Gophersarus runs many different open source packages under the hood.  Again, we did not reinvent the wheel, we built the glue, and then cut and pasted the wheel together. ;)

Please note that the `github.com/mattn/go-sqlite3` package is missing from the list below.  That is because `go install github.com/mattn/go-sqlite3` was executed ahead previously to save time.  The `github.com/mattn/go-sqlite3` package is primarily written in C and therefore takes much longer to compile than pure go code.

Run `go build -v` to see all packages as they build:

```
→ go build -v                                                                                                       

gopkg.in/mgo.v2/bson
gopkg.in/mgo.v2/internal/scram
github.com/asaskevich/govalidator
github.com/codegangsta/negroni
github.com/gorilla/context
github.com/gorilla/mux
github.com/lib/pq/oid
gopkg.in/gorp.v1
github.com/lib/pq
gopkg.in/unrolled/render.v1
gopkg.in/mgo.v2
gopkg.in/validator.v2
gopkg.in/yaml.v2
github.com/codegangsta/cli
git.target.com/gophersaurus/gf.v1/imgo
git.target.com/gophersaurus/gf.v1
git.target.com/gophersaurus/gophersaurus/app/models
git.target.com/gophersaurus/gophersaurus/app/services/api.openweathermap.org/data/2.5
git.target.com/gophersaurus/gophersaurus/config
git.target.com/gophersaurus/gophersaurus/bootstrap
git.target.com/gophersaurus/gophersaurus/app/controllers
git.target.com/gophersaurus/gophersaurus/app
git.target.com/gophersaurus/gophersaurus
```

###Contribution guidelines

* Submit an issue.  
* Send us a pull request. 

Thanks! :)

###License

The Gophersaurus gf.v1 is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT)
