package bootstrap

import (
	"fmt"

	"github.com/gophersaurus/gf.v1/config"
)

// Config bootstraps the initalization and reading of configuration files.
//
// This is the method you would like to edit if your application requires more
// advanced configuration settings like reading from etcd or consul.
//
// import "github.com/gophersaurus/gf.v1/config"
//
// config.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
// config.SetConfigType("yaml") // because there is no file extension in a stream of bytes
//
// // read from remote config the first time.
// err := config.ReadRemoteConfig()
//
// // marshal config
// config.Marshal(&runtime_conf)
//
// // open a goroutine to wath remote changes forever
// go func(){
//     for {
//         time.Sleep(time.Second * 5) // delay after each request
//
//         // currenlty, only tested with etcd support
//         err := config.WatchRemoteConfig()
//         if err != nil {
//             log.Errorf("unable to read remote config: %v", err)
//             continue
//         }
//
//         // marshal new config into our runtime config struct. you can also use channel
//         // to implement a signal to notify the system of the changes
//         config.Marshal(&runtime_conf)
//     }
// }()
//
func Config() error {

	// read configuration file paths
	err := config.ReadInConfig()
	if err != nil {
		return fmt.Errorf("app config settings error: %s \n", err)
	}

	return nil
}
