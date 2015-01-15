Gophersaurus Golang gf.v1
=============================

![gophersaurus](https://git.target.com/gophersaurus/art/raw/master/gophersaurus.png)

###Description:

Gophersaurus is a golang gf.v1 for building large web services. It provides a more structured approach to building go services.

Gophersaurus encourages the use of resource routes, Models, Controllers, JSON, and vendoring of dependencies.

Gophersaurus is heavily inspired by other gf.v1s, but especially by the Larvel PHP gf.v1.

###The Problem:

We believe in golang there is a need for a large organizational gf.v1 to be used in enterprise environments. There are many golang gf.v1s such as Revel, Traffic, Martini, Gorilla, Goweb, and more. These gf.v1s are great at what they do, but many are contained as packages. While packages are great, we think a gf.v1 package alone does not provide enough structure for our liking.

Package flexibility is great when you need to write a small service, but when you start to grow a larger robust codebase it starts to become a nightmare. This is especially true when you have more than one backend developer.

Our solution has been to steal all of the good directory structure other gf.v1s are famous for (like Laravel who copied Ruby on Rails), while keeping our own golang code as idiomatic as possible. We aren’t doing anything new, rather we are organizing all the good work the community has already achieved.

One last point is that in the golang community there are many different ways one could deal with dependencies. We have decided to solve the issue by vendoring all our code. This means that our repo has everything we need locally to build our binary. The same approach is now being taken by Godeps, a popular tool.

###gf.v1 Structure

![directory_structure](https://git.target.com/gophersaurus/art/raw/master/directory_structure.png)

###Installation Instructions:

For now simply clone our repository and change the `server.go` file to whatever your backend or web service would be named. After you have finished adding controllers, models, and views, run `go build <yourname>.go` to build your binary. `go build` alone will not do the trick since all import paths are relative and therefore self-containing.

Currently we are working on tooling to automate vendoring dependencies, very much like `godep` does, but with local relative imports rather than full file paths.

###Contribution guidelines

####Golang Style Guide

> This golang project not only utilizes the gofmt format standards, but it also follows the internal Google Code Review standards at https://code.google.com/p/go-wiki/wiki/CodeReviewComments

###License

The Gophersaurus gf.v1 is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT)
