Gophersaurus
=============================

Gophersaurus is a framework for building monolithic web services quickly. It provides a more structured approach to building go services and abstracts much of the common logic away for building an API.  These traits make Gophersaurus a great choice for projects that need to be rapidly developed, yet maintain consistency.  

![gophersaurus](https://git.target.com/gophersaurus/art/raw/master/gophersaurus.png)

Gophersaurus provides structure for large teams of gophers, and wraps many open source golang packages to abstract away common logic.  It is for these reasons Gophersaurus got its name.  Its a big package.

Gophersaurus has been tested in production, it can scale, and basically gets the job done, but it is not the framework for trying to eek out every bit of performance possible.  As the framework matures, it will become more performant, but honestly... if you want the best golang performance possible, don't use a framework.  The go standard library will do.  

Gophersaurus is heavily inspired by other backend frameworks, but especially the Laravel PHP Framework.  Just like Laravel, Gophersaurus encourages the use of Models, Controllers, Resources, and JSON views/responses.

> IMPORTANT NOTE: Gophersaurus is still in development and currently the API is not stable.  We will lockdown the API soon, and then gf.v1 will cease to have breaking changes.  All breaking changes will be diverted gf.v2.

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

An example YAML configuration file is provided below:

```YAML
config:
  port: 8080
  keys:
    acceptOnlyTheseRefererKey:
    - google.com:80
    - cnn.com:8080
    acceptAnythingDevKey:
  databases:
  - type: mongo
  name: mongoDatabaseName
  user: mongoUserName
  pass: mongoUserPassword
  address: localhost:27017/mongoDatabaseName
```

###The Problem:

We believe in golang there is a need for a large organizational gf.v1 to be used in enterprise environments. There are many golang gf.v1s such as Revel, Traffic, Martini, Gorilla, Goweb, and more. These gf.v1s are great at what they do, but many are contained as packages. While packages are great, we think a gf.v1 package alone does not provide enough structure for our liking.

Package flexibility is great when you need to write a small service, but when you start to grow a larger robust codebase it starts to become a nightmare. This is especially true when you have more than one backend developer.

Our solution has been to steal all of the good directory structure other gf.v1s are famous for (like Laravel who copied Ruby on Rails), while keeping our own golang code as idiomatic as possible. We aren’t doing anything new, rather we are organizing all the good work the community has already achieved.

One last point is that in the golang community there are many different ways one could deal with dependencies. We have decided to solve the issue by vendoring all our code. This means that our repo has everything we need locally to build our binary. The same approach is now being taken by Godeps, a popular tool.

###Installation Instructions:

For now simply clone our repository and change the `server.go` file to whatever your backend or web service would be named. After you have finished adding controllers, models, and views, run `go build <yourname>.go` to build your binary. `go build` alone will not do the trick since all import paths are relative and therefore self-containing.

Currently we are working on tooling to automate vendoring dependencies, very much like `godep` does, but with local relative imports rather than full file paths.

###Contribution guidelines

####Golang Style Guide

> This golang project not only utilizes the gofmt format standards, but it also follows the internal Google Code Review standards at https://github.com/golang/go/wiki/CodeReviewComments

###License

The Gophersaurus gf.v1 is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT)
