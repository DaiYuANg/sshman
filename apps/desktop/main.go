package main

import (
	"github.com/daiyuang/sshman/desktop/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
