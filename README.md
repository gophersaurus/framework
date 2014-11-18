Gophersaurus Golang Framework
=============================

![gophersaurus](https://git.target.com/gophersaurus/art/raw/master/gophersaurus.png)

###Description:

Gophersaurus is a golang framework for building web services. It combines many open source golang libraries into a more structured approach when building go service binaries.

Gophersaurus encourages the use of resource routes, Model-View-Controller, JSON support, and vendoring of dependencies for clean builds.

Gophersaurus is heavily inspired by other golang frameworks and the Larvel PHP Framework.

###The Problem:

We believe in golang there is a need for an organizational framework (especially in large enterprise environments). To be honest there are many golang frameworks out there already and you might have heard of some (Revel, Traffic, Martini, Gorilla, Goweb, and more…). We believe those frameworks don’t provide enough structure for our liking.

Many gophers are skeptical of using frameworks because they fear code bloat and runtime reflection that some of these frameworks (especially Martini) are famous for using. Truth be told, you can indeed use only the “http” package (http://golang.org/pkg/net/http/) for most projects.

The issue we have found is that most other frameworks provide very little structure for your application. In fact all of these frameworks act more like packages that give you a large feature set. This package structure also allows LOTS of flexibility… too much flexibility… The “http” package golang provides a out-of-the-box extremely powerful package, but it is even MORE flexible. So much flexibility is great when you need to write a small service, but when you start to grow a larger robust codebase it starts to become a nightmare. This is especially true when you have more than one backend developer.

Our solution has been to essentially steal all of the good directory structure other frameworks are famous for (like Laravel who copied Ruby on Rails), while keeping our own golang code as idiomatic as possible. We aren’t doing anything “new”, rather we are organizing all the good work the community has already achieved. We don’t use runtime reflection, except for some initialization points like parsing config.json options. Also we are trying to take the best open source go libraries and include them into the framework (gorilla for sessions and route paths or mgo for MongoDB connections are just some examples), until those portions of code get replaced.

One last point is that in the golang community there are many different ways one could deal with dependencies. We have decided to solve the issue by vendoring all our code. This means that our repo has everything we need locally to build our binary. The same approach is now being taken by Godeps, a popular tool.

###Framework Structure

![directory_structure](https://git.target.com/gophersaurus/art/raw/master/directory_structure.png)

###Installation Instructions:

For now simply clone our repository and change the `server.go` file to whatever your backend or web service would be named. After you have finished adding controllers, models, and views, run `go build <yourname>.go` to build your binary. `go build` alone will not do the trick since all import paths are relative and therefore self-containing.

Currently we are working on tooling to automate vendoring dependencies, very much like `godep` does, but with local relative imports rather than full file paths.

###Contribution guidelines

####Golang Style Guide

> This golang project not only utilizes the gofmt format standards, but it also follows the internal Google Code Review standards at https://code.google.com/p/go-wiki/wiki/CodeReviewComments

###License

The Gophersaurus framework is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT)
