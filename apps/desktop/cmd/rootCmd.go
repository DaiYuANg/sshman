package cmd

import (
	"context"
	"github.com/daiyuang/sshman/core"
	"github.com/daiyuang/sshman/desktop/internal"
	"github.com/spf13/cobra"
)

var guiCmd = &cobra.Command{
	Use: "gui",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.CreateContainer(internal.Module).Start(context.Background())
	},
}

func Execute() error {
	return guiCmd.Execute()
}
