Blueprint Golang Framework
==========================

###Description:

Blueprint is a golang framework for building web services. It combines many open source golang libraries into a more structured approach when building go binaries.

Blueprint encourages the use of resource routes, Model-View-Controller, JSON support, and vendoring of dependencies for clean builds.

Blueprint is heavily inspired by other golang frameworks and the Larvel PHP Framework.

###Installation Instructions:

For now simply clone our repository and change our `server.go` file to whatever your backend or web service would be named. After you have finished adding controllers, models, and more, run `go build <yourname>.go` to build your binary. `go build` alone will not do the trick since all import paths are relative and therefore self-containing.

Currently we are working on a tool to automate vendoring dependencies, very much like `godep` does, but with local relative imports rather than full file paths.

###Contribution guidelines

####Golang Style Guide

> This golang project not only utilizes the gofmt format standards, but it also follows the internal Google Code Review standards at https://code.google.com/p/go-wiki/wiki/CodeReviewComments

Some links to look at:

http://stackoverflow.com/questions/24873883/organizing-environment-variables-golang

http://peter.bourgon.org/go-in-production/

###License

The Blueprint framework is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)
