package controllers

import "git.target.com/gophersaurus/gophersaurus/config"

var conf config.Config

// Init takes some parameters to initalizes the controller package variables.
func Init(c config.Config) {
	conf = c
}
