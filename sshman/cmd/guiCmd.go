package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"sman/internal/gui"
)

var guiCmd = &cobra.Command{
	Use: "gui",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createContainer(gui.Module).Start(context.Background())
	},
}
