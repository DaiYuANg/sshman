package cmd

import (
	"context"
	"github.com/daiyuang/sshman/core"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use: "ssh",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.CreateContainer().Start(context.Background())
	},
}

func init() {
	sshCmd.Flags().StringP("port", "p", "22", "Port to connect to on the remote host")
	sshCmd.Flags().StringP("identity-file", "i", "", "Identity (private key) file")
	sshCmd.Flags().StringArrayP("option", "o", nil, "SSH options in format 'Key=Value'")
	sshCmd.Flags().BoolP("verbose", "v", false, "Verbose mode (use -vvv for more verbosity)")
	sshCmd.Flags().Bool("agent-forwarding", false, "Enable SSH agent forwarding")
	sshCmd.Flags().Bool("x11-forwarding", false, "Enable X11 forwarding")
	sshCmd.Flags().StringArrayP("local-forward", "L", nil, "Local port forwarding (e.g. 8080:remotehost:80)")
	sshCmd.Flags().StringArrayP("remote-forward", "R", nil, "Remote port forwarding (e.g. 8080:localhost:80)")
	sshCmd.Flags().Bool("compression", false, "Enable compression")
	sshCmd.Flags().Duration("connect-timeout", 0, "Timeout for establishing connection (e.g. 10s)")
	sshCmd.Flags().Bool("no-host-key-check", false, "Disable strict host key checking")
	sshCmd.Flags().String("user", "", "Username for SSH connection (overrides user@host)")
}
