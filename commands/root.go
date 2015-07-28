package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "Gophersaurus App",
	Short: "Gophersaurus is a RESTful API framework for Golang services.",
	Long: `Gophersaurus is a RESTFul API framework
          that provides a solid directory structure
          to organize a MVC application.`,
}
