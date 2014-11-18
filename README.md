Gophersaurus Golang Framework
=============================

![gophersaurus](https://git.target.com/gophersaurus/art/blob/master/gophersaurus.png)

###Description:

Gophersaurus is a golang framework for building web services. It combines many open source golang libraries into a more structured approach when building go service binaries.

Gophersaurus encourages the use of resource routes, Model-View-Controller, JSON support, and vendoring of dependencies for clean builds.

Gophersaurus is heavily inspired by other golang frameworks and the Larvel PHP Framework.

###Installation Instructions:

For now simply clone our repository and change the `server.go` file to whatever your backend or web service would be named. After you have finished adding controllers, models, and views, run `go build <yourname>.go` to build your binary. `go build` alone will not do the trick since all import paths are relative and therefore self-containing.

Currently we are working on tooling to automate vendoring dependencies, very much like `godep` does, but with local relative imports rather than full file paths.

###Contribution guidelines

####Golang Style Guide

> This golang project not only utilizes the gofmt format standards, but it also follows the internal Google Code Review standards at https://code.google.com/p/go-wiki/wiki/CodeReviewComments

###License

The Gophersaurus framework is open-sourced software licensed under the [MIT License](http://opensource.org/licenses/MIT)
