package main

import (
	cmd "github.com/octoberswimmer/apexfmt/cmd"

	"github.com/spf13/cobra/doc"
)

func main() {
	cmd.RootCmd.DisableAutoGenTag = true
	err := doc.GenMarkdownTree(cmd.RootCmd, "./docs")
	if err != nil {
		panic(err.Error())
	}
}
