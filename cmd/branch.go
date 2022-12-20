package cmd

import "github.com/spf13/cobra"

var branch = &cobra.Command{
	Use:   "branch [branch name]",
	Short: "Branches can be switched by the default parameter",
	Long:  "If you want some further operations, it is a good way to use extra parameters",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
