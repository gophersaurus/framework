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
1. Clone this repositiory to your local machine. `git clone git@git.target.com:gophersaurus/gophersaurus.git`
2. Enter the root project directory. `cd gophersaurus`
4. Run the command `go get ./...` to ensure you have all nessesary dependencies locally.  If you want download fresh updated dependency copies and see what what your downloading run `go get -u -v ./...`.
5. Rename the `gophersaurus` to your project name. `cd ..`, `mv gophersaurus project-name`
6. Do a search and replace in your project root for `git.target.com/gophersaurus/gophersaurus` to `your.git.com/your-org/your-project-name`

####Git Setup
7. In the project directory run `git remote set-url origin git@your.git.com:your-org/your-project-name.git`

Now you should be able to run `go build` from your project root to build and all manner of `git` commands.

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

> Note: Referer `localhost` values currently translate as `::1`.  Most `/etc/hosts` files have `::1` listed last after `127.0.0.1`.  Also do not attempt to compensate for proxies or loadbalancers unless you know what your doing.  Gophersaurus will search the HTTP `Header` for a `X-FORWARDED-FOR` value by default.

###Style Guide

We believe it is important for a framework to provide a style guide, not just code.  

Instead of reinventing the wheel, we recommend `gofmt` and `goimports` to automatically format go code properly.  Beyond these awesome tools we also recommend gophers to keep close to the internal Golang Code Review standards at https://github.com/golang/go/wiki/CodeReviewComments.

###What Problem Is Gophersaurus Solving?

We believe in golang there is a need for a large framework to be used in enterprise environments. There are many golang frameworks such as Revel, Traffic, Martini, Gorilla, Goweb, and more. These frameworks are great at what they do, but lack folder structure and strong opinions.

Package flexibility is great when you need to write a small service, but when you start to grow a larger robust codebase it starts to become a nightmare. This is especially true when you have more than one developer.

Our solution/plan has been to steal all of the good directory structure other frameworks are famous for (like Laravel who copied Ruby on Rails), while keeping our own golang code as idiomatic as possible. We aren’t doing anything new, rather we are organizing all the good work the community has already achieved.

###Contribution guidelines

###License

The Gophersaurus gf.v1 is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT)
