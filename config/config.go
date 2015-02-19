package config

import (
	"log"

	"git.target.com/gophersaurus/gf.v1"
)

// Config describes configuration settings.
//
// Please note that we do not recommend splitting settings by enviroment.
// Doing this encourages one large file that contains all the sensative
// information for dev, stage, prod, etc...
//
// The project .gitignore attempts to stop you from checking your config file
// into your version control system.  Even if your project is checked into a
// private repository, we recommend you do not include your config file.
//
// You might disagree with us, but this is for your own good.
// Your security friends will thank us.
//
type Config struct {
	Env string
	gf.Config
}

// ReadFile takes a filename and returns a Config object.
func ReadFile(filename string) Config {

	// Create a new Config object to work with.
	config := Config{}

	// Read the file values into the Config object.
	if err := gf.ReadConfig(filename, config); err != nil {

		// If we have an error, we should log the error and exit.
		// Invalid configuration is a show stopper.
		log.Fatal(err)
	}

	return config
}
