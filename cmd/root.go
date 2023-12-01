package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/formatter"
	"github.com/octoberswimmer/apexfmt/parser"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func init() {
	cobra.OnInitialize(globalConfig)
	RootCmd.Flags().BoolP("write", "w", false, "write result to (source) file instead of stdout")
}

var RootCmd = &cobra.Command{
	Use:   "apexfmt [file...]",
	Short: "Format Apex",
	Run: func(cmd *cobra.Command, args []string) {
		write, _ := cmd.Flags().GetBool("write")
		for _, filename := range args {
			src, err := readFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read file: %s\n", err.Error())
				os.Exit(1)
			}
			out, err := format(string(src))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to format file %s: %s\n", filename, err.Error())
				os.Exit(1)
			}

			if !write {
				fmt.Println(out)
			}
			if write && string(src) != out {
				err = writeFile(filename, []byte(out))
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

func format(src string) (string, error) {
	input := antlr.NewInputStream(src)
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewApexParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	v := formatter.NewVisitor(stream)
	out, ok := p.CompilationUnit().Accept(v).(string)
	if !ok {
		return "", fmt.Errorf("Unexpected result parsing apex")
	}
	return out, nil
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	src, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func writeFile(filename string, contents []byte) error {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %s\n", err.Error())
		os.Exit(1)
	}
	perm := info.Mode().Perm()
	size := info.Size()
	fout, err := os.OpenFile(filename, os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	defer fout.Close()
	n, err := fout.Write(contents)
	if err == nil && int64(n) < size {
		err = fout.Truncate(int64(n))
	}
	return err
}
