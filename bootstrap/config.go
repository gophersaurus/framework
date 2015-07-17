package bootstrap

import (
	"log"

	"github.com/gophersaurus/framework/config"
)

// Config takes settings and returns a Config object.
func Config(settings map[string]string) config.Config {

	// The first thing we need to initalize is configuration settings.
	c := config.New()

	// Check if the config file has been set.
	if conf, ok := settings["config"]; ok && len(conf) > 0 {

		// Read the defined config file.
		if err := config.Read(conf, c); err != nil {
			log.Fatalln(err)
		}

	} else {

		// Assume the name of the configuration file.
		if err := config.Read("config/config.yml", c); err != nil {
			log.Fatalln(err)
		}
	}

	return c
}
