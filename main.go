package main

import "github.com/omarisadev/github-activity/cmd"

func main() {
	rootCmd := cmd.RootCmd()
	rootCmd.Execute()
}
