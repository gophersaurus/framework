package main

import (
	"os"

	"git.target.com/gophersaurus/gophersaurus/bootstrap"
	"github.com/codegangsta/cli"
)

// The go program starts here.
func main() {

	// Get all the flag stuff you need here.
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.json",
			Usage: "The path to the config file. Defaults to 'config.json' in the execution path.",
		},
		cli.StringFlag{
			Name:  "env, e",
			Value: "dev",
			Usage: "The name of the environment to be used",
		},
	}
	app.Action = func(c *cli.Context) {
		// get parameters from flags
		path := c.String("config")
		env := c.String("env")

		// Start serving the application.
		bootstrap.Init(path, env)
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
