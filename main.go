package main

import "github.com/gophersaurus/framework/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.ServeCmd)
	cmd.RootCmd.Execute()
}
