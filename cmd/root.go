package cmd

import (
	"fmt"
	"os"

	"github.com/octoberswimmer/apexfmt/formatter"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func init() {
	cobra.OnInitialize(globalConfig)
	RootCmd.Flags().BoolP("write", "w", false, "write result to (source) file instead of stdout")
	RootCmd.Flags().BoolP("list", "l", false, "list files whose formatting differs from apexfmt's")
}

var RootCmd = &cobra.Command{
	Use:   "apexfmt [file...]",
	Short: "Format Apex",
	Run: func(cmd *cobra.Command, args []string) {
		write, _ := cmd.Flags().GetBool("write")
		list, _ := cmd.Flags().GetBool("list")
		for _, filename := range args {
			f := formatter.NewFormatter(filename)
			err := f.Format()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to format file %s: %s\n", filename, err.Error())
				os.Exit(1)
			}

			if list {
				changed, err := f.Changed()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to check file for changes %s: %s\n", filename, err.Error())
					os.Exit(1)
				}
				if changed {
					fmt.Println(filename)
				}
			} else if !write {
				out, err := f.Formatted()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to get formatted source %s: %s\n", filename, err.Error())
					os.Exit(1)
				}
				fmt.Println(out)
			}
			changed, err := f.Changed()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to check file for changes %s: %s\n", filename, err.Error())
				os.Exit(1)
			}
			if write && changed {
				err = f.Write()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to write file %s: %s\n", filename, err.Error())
					os.Exit(1)
				}
			}
		}
	},
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
}

func globalConfig() {
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
