// Package config manages configuration settings for the gophersaurus framework.
// In reality, config is really just a wrapper around the viper package.
//
// You can learn more about viper at https://github.com/spf13/viper.
package config

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {

	// env/config path locations
	SetConfigName(".env")
	AddConfigPath("./")

	// defaults
	SetDefault("port", "5225")
}

// Public vars
var (
	RemoteConfig             = viper.RemoteConfig
	SupportedExts            = viper.SupportedExts
	SupportedRemoteProviders = viper.SupportedRemoteProviders
)

// RemoteConfigError denotes encountering an error while trying to
// pull the configuration from the remote provider.
type RemoteConfigError string

// Returns the formatted remote provider error
func (rce RemoteConfigError) Error() string {
	return fmt.Sprintf("Remote Configurations Error: %s", string(rce))
}

// RemoteProvider stores the configuration necessary
// to connect to a remote key/value store.
// Optional secretKeyring to unencrypt encrypted values
// can be provided.
type RemoteProvider interface {
	Provider() string
	Endpoint() string
	Path() string
	SecretKeyring() string
}

// UnsupportedConfigError denotes encountering an unsupported
// configuration filetype.
type UnsupportedConfigError string

// Returns the formatted configuration error.
func (str UnsupportedConfigError) Error() string {
	return fmt.Sprintf("Unsupported Config Type %q", string(str))
}

// UnsupportedRemoteProviderError denotes encountering an unsupported remote
// provider. Currently only Etcd and Consul are
// supported.
type UnsupportedRemoteProviderError string

// Returns the formatted remote provider error.
func (str UnsupportedRemoteProviderError) Error() string {
	return fmt.Sprintf("Unsupported Remote Provider Type %q", string(str))
}

// AddConfigPath adds a path for Viper to search for the config file in.
// Can be called multiple times to define multiple search paths.
func AddConfigPath(in string) {
	viper.AddConfigPath(in)
}

// AddRemoteProvider adds a remote configuration source.
// Remote Providers are searched in the order they are added.
// Provider is a string value, "etcd" or "consul" are currently supported.
// Endpoint is the url. etcd requires http://ip:port consul requires ip:port
// path is the path in the k/v store to retrieve configuration.
// To retrieve a config file called myapp.json from /configs/myapp.json you
// should set path to /configs and set config name (SetConfigName()) to "myapp"
func AddRemoteProvider(provider, endpoint, path string) error {
	return viper.AddRemoteProvider(provider, endpoint, path)
}

// AddSecureRemoteProvider adds a remote configuration source.
// Secure Remote Providers are searched in the order they are added.
// Provider is a string value, "etcd" or "consul" are currently supported.
// Endpoint is the url. etcd requires http://ip:port consul requires ip:port
// secretkeyring is the filepath to your openpgp secret keyring. e.g.
// /etc/secrets/myring.gpg path is the path in the k/v store to retrieve
// configuration To retrieve a config file called myapp.json from
// /configs/myapp.json you should set path to /configs and set config name
// (SetConfigName()) to "myapp" Secure Remote Providers are implemented with
// github.com/xordataexchange/crypt
func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error {
	return viper.AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
}

// AllKeys returns all keys regardless where they are set.
func AllKeys() []string {
	return viper.AllKeys()
}

// AllSettings returns all settings as a map[string]interface{}.
func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

// AutomaticEnv checks ENV variables for all keys set in config, default &
// flags.
func AutomaticEnv() {
	viper.AutomaticEnv()
}

// BindEnv binds a Viper key to a ENV variable.
//
// ENV variables are case sensitive.
//
// If only a key is provided, BindEnv will use the ENV key matching the key,
// uppercased. EnvPrefix will be used when set when ENV name is not provided.
func BindEnv(input ...string) (err error) {
	return viper.BindEnv(input...)
}

// BindPFlag binds a specific key to a flag (as used by cobra).
// Example(where serverCmd is a Cobra instance):
//
//	 serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
//	 Viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
//
func BindPFlag(key string, flag *pflag.Flag) (err error) {
	return viper.BindPFlag(key, flag)
}

// BindPFlags binds a full flag set to the configuration, using each flag's long
// name as the config key.
func BindPFlags(flags *pflag.FlagSet) (err error) {
	return viper.BindPFlags(flags)
}

// FileUsed returns the file used to populate the config registry.
func FileUsed() string {
	return viper.ConfigFileUsed()
}

// Debug prints all configuration registries for debugging purposes.
func Debug() {
	viper.Debug()
}

// Get can retrieve any value given the key to use
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) interface{} {
	return viper.Get(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetSizeInBytes returns the size of the value associated with the given key
// in bytes.
func GetSizeInBytes(key string) uint {
	return viper.GetSizeInBytes(key)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetStringMap returns the value associated with the key as a map of
// interfaces.
func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

// GetStringMapString returns the value associated with the key as a map of
// strings.
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// GetStringMapStringSlice returns the value associated with the key as a map of
// a slice of strings.
// func GetStringMapStringSlice(key string) map[string][]string {
// return viper.GetStringMapStringSlice(key)
// }
func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

// GetStringSlice returns the value associated with the key as a slice of
// strings.
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

// InConfig checks to see if the given key (or an alias) is in the config file.
func InConfig(key string) bool {
	return viper.InConfig(key)
}

// IsSet checks to see if the key has been set in any of the data locations.
func IsSet(key string) bool {
	return viper.IsSet(key)
}

// Unmarshal unmarshals the config into a Struct. Make sure that the tags
// on the fields of the structure are properly set.
func Unmarshal(rawVal interface{}) error {
	return viper.Unmarshal(rawVal)
}

// UnmarshalKey takes a single key and unmarshal it into a Struct.
func UnmarshalKey(key string, rawVal interface{}) error {
	return viper.UnmarshalKey(key, rawVal)
}

// ReadConfig reads the configuration of an io.Reader.
func ReadConfig(in io.Reader) error {
	return viper.ReadConfig(in)
}

// ReadInConfig will discover and load the configuration file from disk
// and key/value stores, searching in one of the defined paths.
func ReadInConfig() error {
	return viper.ReadInConfig()
}

// ReadInEnvConfig reads the configuration of an io.Reader and will map
// environment key values to the specific env value provided when calling:
//
// framework serve --env prod
//
// An example of an io.Reader that would contain environment specific key values
// might be:
//
// environments:
// 		prod:
// 				KEY: value
//
func ReadInEnvConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	env := viper.GetString("env")
	if len(env) > 0 {

		environments := viper.GetStringMapString("environments." + env)
		if len(environments) < 1 {
			return fmt.Errorf("environment '%s' not found in config file", env)
		}

		for key, value := range environments {
			viper.Set(key, value)
		}
	}
	return nil
}

// ReadRemoteConfig attempts to get configuration from a remote source
// and read it in the remote configuration registry.
func ReadRemoteConfig() error {
	return viper.ReadRemoteConfig()
}

// RegisterAlias provide another accessor for the same key.
// This enables one to change a name without breaking the application.
func RegisterAlias(alias string, key string) {
	viper.RegisterAlias(alias, key)
}

// Reset is intended for testing, will reset all to default settings.
// In the public interface for the viper package so applications
// can use it in their testing as well.
func Reset() {
	viper.Reset()
}

// Set sets the value for the key in the override regiser.
// Will be used instead of values obtained via
// flags, config file, ENV, default, or key/value store.
func Set(key string, value interface{}) {
	viper.Set(key, value)
}

// SetConfigFile explicitly define the path, name and extension of the config
// file to use and not check any of the config paths
func SetConfigFile(in string) {
	viper.SetConfigFile(in)
}

// SetConfigName names the config file, but does not include the extension.
func SetConfigName(in string) {
	viper.SetConfigName(in)
}

// SetConfigType sets the type of the configuration returned by the
// remote source, e.g. "json".
func SetConfigType(in string) {
	viper.SetConfigType(in)
}

// SetDefault sets the default value for this key.
// Default only used when no value is provided by the user via flag, config or
// ENV.
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// SetEnvKeyReplacer sets the strings.Replacer on the viper object
// Useful for mapping an environmental variable to a key that does
// not match it.
func SetEnvKeyReplacer(r *strings.Replacer) {
	viper.SetEnvKeyReplacer(r)
}

// SetEnvPrefix defines a prefix that ENVIRONMENT variables will use.
// E.g. if your prefix is "spf", the env registry
// will look for ENV variables that start with "SPF_".
func SetEnvPrefix(in string) {
	viper.SetEnvPrefix(in)
}

// WatchRemoteConfig watches a remote configuration source such as etcd.
func WatchRemoteConfig() error {
	return viper.WatchRemoteConfig()
}
