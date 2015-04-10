package controllers

import "github.com/gophersaurus/framework/config"

var conf config.Config

// Init takes some parameters to initalizes the controller package variables.
func Init(c config.Config) {
	conf = c
}
