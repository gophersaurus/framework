package main

import (
	"fmt"

	"git.target.com/gophersaurus/gophersaurus/app"
	"github.com/codegangsta/cli"
)

// main starts the program.
func main() {

	// Create a new CLI application.
	app := cli.NewApp()

	// Define default CLI flags.
	app.Flags = []cli.Flag{

		// Configuration options.
		cli.StringFlag{
			Name:  "config, conf, c",
			Value: "config.yml",
			Usage: "The path to the config file. Defaults to 'config.yml'.",
		},

		// Environment options.
		cli.StringFlag{
			Name:  "env, e",
			Value: "dev",
			Usage: "The application environment.",
		},

		// Static file path.
		cli.StringFlag{
			Name:  "static, s",
			Value: "/public",
			Usage: "The environment for the application to run in.",
		},
	}

	// Define the default CLI action.
	app.Action = func(c *cli.Context) {

		// Start with nice spacing :)
		fmt.Print("\n")

		// Define the application settings.
		settings := map[string]string{
			"config": c.String("config"),
			"env":    c.String("env"),
			"static": c.String("static"),
		}

		// Bootstrap the server.
		s := server.Bootstrap(settings)

		// Start serving content.
		s.Serve()
	}

	// Run the CLI application.
	app.RunAndExitOnError()
}
