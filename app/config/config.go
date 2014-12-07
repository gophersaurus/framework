package config

import (
	"errors"
	"fmt"

	gf "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/framework"
)

var Env ExpandedEnvConfig

type ApiConfig map[string]ExpandedEnvConfig

type ExpandedEnvConfig struct {
	gf.EnvConfig
}

func Init(path, env string) {

	path, err := gf.GetValidConfigPath(path)
	fmt.Println("checking config path exists...")
	gf.Check(err)

	var conf ApiConfig
	err = gf.LoadConfig(path, &conf)
	fmt.Println("checking config loaded properly...")
	gf.Check(err)

	envObj, exists := conf[env]
	if !exists {
		gf.Check(errors.New("environment settings for " + env + " were not found in config json file"))
	}
	Env = envObj
}
