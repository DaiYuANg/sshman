package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "sman",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
	rootCmd.AddCommand(sshCmd)
	rootCmd.AddCommand(guiCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
