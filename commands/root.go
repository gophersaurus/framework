package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "gophersaurus application",
	Short: "gophersaurus is a RESTful API framework for building go services.",
	Long: `gophersaurus is a RESTFul API framework
          that provides a solid structure for developing and organizing MVC
					application services.`,
}
