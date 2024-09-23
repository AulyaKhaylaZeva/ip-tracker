package cmd

import (
	"github.com/spf13/cobra"
)

var version = "1.0.0"

var (
	rootCmd = &cobra.Command{
		Use:     "iptracker",
		Version: version,
		Short:   "CLI App to track IP Addresses.",
		Long:    `CLI App to track IP Addresses.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
