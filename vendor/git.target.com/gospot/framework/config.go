package framework

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// The Config struct that holds all the config settings.
type Config map[Env]EnvConfig

var ConfigPath string

type Env string

type EnvConfig struct {
	App AppConfig
	DB  *DbConfig
}

type AppConfig struct {
	Port       int
	IndentJson bool
	Keys       KeyMap
}

type KeyMap map[Key]KeyConfig

func (k *KeyMap) Get(key string) *KeyConfig {
	conf, exists := (*k)[Key(key)]
	if !exists {
		return nil
	}
	return &conf
}

type Key string

type KeyConfig struct {
	Status bool
	Urls   []string // whitelist url
}

type DbConfig struct {
	Type     string
	Addr     string
	Name     string
	Username string
	Password string
}

func (db *DbConfig) IsValid() bool {
	return db.Type == "mongo" && len(db.Addr) > 0 && len(db.Username) > 0 && len(db.Password) > 0
}

func GetValidConfigPath(path string) (string, error) {
	dir, file := filepath.Split(path)

	if len(path) < 1 {
		dir = filepath.Dir(os.Args[0])
		file = "/config.json"
	}

	_, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	return dir + file, nil
}

func LoadConfig(path string, c interface{}) error {
	// Read the config.json file into a slice of bytes.
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	ConfigPath = path

	// Build the config object.
	err = json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetEnvironment(env Env) (*EnvConfig, error) {
	base, ok := (*c)[env]
	if ok {
		return &base, nil
	} else {
		return nil, errors.New("environment settings for " + string(env) + " were not found in config json file")
	}
}
