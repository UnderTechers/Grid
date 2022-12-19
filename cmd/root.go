package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grid",
	Short: "Grid - a decentralized version control system with widely supportive extension system",
	Long: `A fast, decentralized and flexible VCS(version control system), along with an extension system to be supportively visual changes in any projects
	`,
}

var echoTimes int

func Execute() {
	flags()
	cmdPrint.Flags().IntVarP(&echoTimes, "times", "", 1, "times to echo the input")
	rootCmd.AddCommand(Init)
	cmdPrint.AddCommand(cmdPrint2)
	rootCmd.AddCommand(cmdPrint)
	rootCmd.AddCommand(sync)
	rootCmd.AddCommand(add)
	rootCmd.AddCommand(listen)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
