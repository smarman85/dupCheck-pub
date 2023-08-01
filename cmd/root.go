package cmd

import (
	"github.com/spf13/cobra"
)

var (
	application string
	environment string
)

func init() {

	// migrate nonsecrets
	rootCmd.AddCommand(check)
	check.Flags().StringVarP(&application, "application", "a", "", " app to filter")
	check.Flags().StringVarP(&environment, "environment", "e", "", "environment to filter")
	check.MarkFlagRequired("application")
	check.MarkFlagRequired("environment")

}

var rootCmd = &cobra.Command{
	Use:   "check",
	Short: "Helper script to check for duplicate values",
}

func Execute() error {
	return rootCmd.Execute()
}
