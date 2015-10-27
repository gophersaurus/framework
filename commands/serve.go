package commands

import (
	"log"

	"github.com/gophersaurus/framework/app"
	"github.com/gophersaurus/gf.v1/config"
	"github.com/spf13/cobra"
)

func init() {

	// bind port flag
	ServeCmd.Flags().IntP("port", "p", 5225, "Port to run Application server on")
	config.BindPFlag("port", ServeCmd.Flags().Lookup("port"))

	// bind static flag
	ServeCmd.Flags().StringP("static", "s", "public", "Where the public static files are")
	config.BindPFlag("static", ServeCmd.Flags().Lookup("static"))

	// bind env flag
	ServeCmd.Flags().StringP("env", "e", "", "The environment that we are running")
	config.BindPFlag("env", ServeCmd.Flags().Lookup("env"))
}

// ServeCmd describes the serve command.
var ServeCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"server", "s"},
	Short:   "Listen and Serve HTTP",
	Long:    "Listen and Serve HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(app.Serve())
	},
}
