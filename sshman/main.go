package main

import (
	"github.com/spf13/cobra"
	"sman/cmd"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
