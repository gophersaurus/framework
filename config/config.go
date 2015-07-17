package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

// Config describes application configuration settings.
//
// Please note that we do not recommend splitting settings by enviroment.
// Doing this encourages one large file that contains all the sensative
// information for dev, stage, prod, etc...
//
// The project .gitignore attempts to stop you from checking your config file
// into your version control system.  Even if your project is checked into a
// private repository, we recommend you do not include the config/config.yml file.
//
// You might disagree with us on this, but it is for your own good.
// Your security friends will thank us.
type Config struct {
	Port        string
	Keys        map[string][]string `yaml:"keys,omitempty",json:"keys,omitempty"`
	Databases   []Database          `yaml:"databases,omitempty",json:"databases,omitempty"`
	TLS         TLS                 `yaml:"tls,omitempty",json:"tls,omitempty"`
	SessionDays int                 `yaml:"session_days,omitempty" json:"session_days,omitempty"`
	Services    `yaml:"services,omitempty" json:"services,omitempty"`
}

// TLS represents cert.pem and key.pem file locations.
type TLS struct {
	Cert string `yaml:"cert,omitempty",json:"cert,omitempty"`
	Key  string `yaml:"key,omitempty",json:"key,omitempty"`
}

// Database represents generic database connection values.
type Database struct {
	Type    string
	Name    string
	User    string
	Pass    string
	Address string
}

// Services represents all third party services.
type Services struct {
	Rackspace `yaml:"rackspace,omitempty" json:"rackspace,omitempty"`
}

// Rackspace represents all the data for a rackspace connection.
type Rackspace struct {
	Key       string `yaml:"key,omitempty" json:"key,omitempty"`
	User      string `yaml:"user,omitempty" json:"user,omitempty"`
	Pass      string `yaml:"pass,omitempty" json:"pass,omitempty"`
	Region    string `yaml:"region,omitempty" json:"region,omitempty"`
	TenantID  string `yaml:"tenantid,omitempty" json:"tenantid,omitempty"`
	Container string `yaml:"container,omitempty" json:"container,omitempty"`
}

// New returns an empty Config.
func New() Config {
	return Config{}
}

// Read takes a file and reads its contents into a config object.
func Read(filename string, config interface{}) error {

	// read config file into bytes
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// unmarshal by file extension
	switch filepath.Ext(filename) {

	// unmarshal json
	case ".json":
		if err := json.Unmarshal(bytes, config); err != nil {
			return err
		}

	// unmarshal yml
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(bytes, config); err != nil {
			return err
		}

	// default on unknown extension types
	default:
		return fmt.Errorf("config file type unsupported: %s", filepath.Ext(filename))
	}

	return nil
}
