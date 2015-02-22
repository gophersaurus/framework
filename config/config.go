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
type Config struct {
	gf.Config
	SessionDays int `yaml:"session_days,omitempty" json:"session_days,omitempty"`
	Services    `yaml:"services,omitempty" json:"services,omitempty"`
}

type Services struct {
	Rackspace `yaml:"rackspace,omitempty" json:"rackspace,omitempty"`
}

type Rackspace struct {
	Key                 string `yaml:"key,omitempty" json:"key,omitempty"`
	User                string `yaml:"user,omitempty" json:"user,omitempty"`
	Pass                string `yaml:"pass,omitempty" json:"pass,omitempty"`
	Region              string `yaml:"region,omitempty" json:"region,omitempty"`
	TenantID            string `yaml:"tenantid,omitempty" json:"tenantid,omitempty"`
	ImageContainer      string `yaml:"imagecontainer,omitempty" json:"imagecontainer,omitempty"`
	PixelcryptContainer string `yaml:"pixelcryptcontainer,omitempty" json:"pixelcryptcontainer,omitempty"`
}

// NewConfig comments
func NewConfig() Config {
	return Config{Config: gf.NewConfig()}
}

// ReadFile takes a filename and returns a Config object.
func ReadFile(filename string) Config {

	c := NewConfig()

	// Read the file values into the Config object.
	if err := gf.ReadConfig(filename, &c); err != nil {

		// If we have an error, we should log the error and exit.
		// Invalid configuration is a show stopper.
		log.Fatal(err)
	}

	return c
}
