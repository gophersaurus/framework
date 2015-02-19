package controllers

import "git.target.com/gophersaurus/gophersaurus/config"

var conf config.Config

func Init(c config.Config) {
	conf = c
}
