package bootstrap

import "github.com/gophersaurus/framework/config"

// Config takes settings and returns a Config object.
func Config(settings map[string]string) config.Config {

	// The first thing we need to initalize is configuration settings.
	c := config.NewConfig()

	// Check if the config file has been set.
	if conf, ok := settings["config"]; ok && len(conf) > 0 {

		// Read the defined config file.
		c = config.ReadFile(conf)

	} else {
		// Assume the name of the configuration file.
		c = config.ReadFile(".config.yml")
	}

	return c
}
