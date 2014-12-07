package bootstrap

import (
	"git.target.com/gophersaurus/gophersaurus/app"
	"git.target.com/gophersaurus/gophersaurus/app/config"
)

// Run the app.
func Init(path, env string) {

	config.Init(path, env)

	db := config.Env.DB
	port := config.Env.App.Port
	indentJson := config.Env.App.IndentJson
	keys := config.Env.App.Keys

	app.Start(keys, db, port, indentJson)
}
