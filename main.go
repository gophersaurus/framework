package main

import "github.com/gophersaurus/framework/commands"

func main() {
	commands.RootCmd.AddCommand(commands.ServeCmd)
	commands.RootCmd.Execute()
}
