package main

import "github.com/gophersaurus/framework/commands"

// main starts the program.
func main() {
	commands.RootCmd.AddCommand(commands.ServeCmd)
	commands.RootCmd.Execute()
}
