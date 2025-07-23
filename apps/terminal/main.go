package main

import (
	"github.com/daiyuang/sshman/terminal/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
