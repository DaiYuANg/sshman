package cmd

import (
	"context"
	"github.com/daiyuang/sshman/core"
	"github.com/daiyuang/sshman/terminal/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "tui",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.CreateContainer(internal.Module).Start(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
