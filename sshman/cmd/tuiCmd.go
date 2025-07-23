package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"sman/internal/tui"
)

var tuiCmd = &cobra.Command{
	Use: "tui",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createContainer(tui.Module).Start(context.Background())
	},
}
