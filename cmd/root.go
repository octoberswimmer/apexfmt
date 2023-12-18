package cmd

import (
	"fmt"
	"os"

	"github.com/octoberswimmer/apexfmt/formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func init() {
	cobra.OnInitialize(globalConfig)
	RootCmd.Flags().BoolP("write", "w", false, "write result to (source) file instead of stdout")
	RootCmd.Flags().BoolP("list", "l", false, "list files whose formatting differs from apexfmt's")
	RootCmd.Flags().BoolP("verbose", "v", false, "enable debug logging")
	RootCmd.Flags().BoolP("soql", "s", false, "format SOQL query")

	RootCmd.MarkFlagsMutuallyExclusive("write", "list")
	RootCmd.MarkFlagsMutuallyExclusive("soql", "write")
	RootCmd.MarkFlagsMutuallyExclusive("soql", "list")

}

var RootCmd = &cobra.Command{
	Use:   "apexfmt [file...]",
	Short: "Format Apex",
	RunE: func(cmd *cobra.Command, args []string) error {
		if soql, _ := cmd.Flags().GetBool("soql"); soql {
			formatSOQL()
			return nil
		}

		write, _ := cmd.Flags().GetBool("write")
		list, _ := cmd.Flags().GetBool("list")
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			log.SetLevel(log.DebugLevel)
		}
		formatters := []*formatter.Formatter{}
		for _, filename := range args {
			formatters = append(formatters, formatter.NewFormatter(filename, nil))
		}
		if len(args) == 0 {
			if write {
				return fmt.Errorf("One or more files required for --write option")
			}
			if list {
				return fmt.Errorf("One or more files required for --list option")
			}
			formatters = append(formatters, formatter.NewFormatter("", os.Stdin))
		}
		for _, f := range formatters {
			err := f.Format()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to format file %s: %s\n", f.SourceName(), err.Error())
				os.Exit(1)
			}

			if list {
				changed, err := f.Changed()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to check file for changes %s: %s\n", f.SourceName(), err.Error())
					os.Exit(1)
				}
				if changed {
					fmt.Println(f.SourceName())
				}
			} else if !write {
				out, err := f.Formatted()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to get formatted source %s: %s\n", f.SourceName(), err.Error())
					os.Exit(1)
				}
				fmt.Println(out)
			}
			changed, err := f.Changed()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to check file for changes %s: %s\n", f.SourceName(), err.Error())
				os.Exit(1)
			}
			if write && changed {
				err = f.Write()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to write file %s: %s\n", f.SourceName(), err.Error())
					os.Exit(1)
				}
			}
		}
		return nil
	},
	DisableFlagsInUseLine: true,
}

func globalConfig() {
}

func formatSOQL() {
	f := formatter.NewSOQLFormatter()
	err := f.Format()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format query: %s\n", err.Error())
		os.Exit(1)
	}
	out, err := f.Formatted()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get formatted query: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(out)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
