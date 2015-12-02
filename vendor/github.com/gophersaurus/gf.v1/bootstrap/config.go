package bootstrap

import (
	"fmt"
	"strings"

	"github.com/gophersaurus/gf.v1/config"
)

// Config bootstraps configuration and environment values.
func Config() error {

	// load configuration settings
	if err := config.ReadInEnvConfig(); err != nil {
		if !strings.Contains(err.Error(), "Unsupported Config Type") {
			return fmt.Errorf("app config settings error: %s \n", err)
		}
	}

	return nil
}
