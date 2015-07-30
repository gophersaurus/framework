package commands

import (
	"github.com/gophersaurus/framework/app"
	"github.com/gophersaurus/gf.v1/config"
	"github.com/spf13/cobra"
)

func init() {

	// bind port flag
	ServeCmd.Flags().Int("port", 5225, "Port to run Application server on")
	config.BindPFlag("port", ServeCmd.Flags().Lookup("port"))

	// bind static flag
	ServeCmd.Flags().String("static", "public", "Where the public static files are")
	config.BindPFlag("static", ServeCmd.Flags().Lookup("static"))
}

// ServeCmd describes the serve command.
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Listen and Serve HTTP",
	Long:  "Listen and Serve HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		app.Serve()
	},
}
