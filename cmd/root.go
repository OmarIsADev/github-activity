package cmd

import (
	"errors"

	"github.com/omarisadev/github-activity/activity"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github-activity",
		Short: "Github User Activity is a CLI tool for fetching user activity",
		Long: `Github User Activity is a CLI tool for fetching user activity. It allows you to fetch user activity by providing the username.

Example:
> github-activity omarisadev

Complete code available at https://github.com/omarisadev/github-activity`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDisplayActivityCmd(args)
		},
	}

	return cmd
}

func RunDisplayActivityCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("username is required")
	}

	return activity.GetActivity(args[0])
}
