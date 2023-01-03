package cmd

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/tidwall/sjson"
)

func init() {
	rootCmd.AddCommand(branch)
}

var branch = &cobra.Command{
	Use:   "branch [branch name]",
	Short: "Branches can be switched by the default parameter",
	Long:  "If you want some further operations, it is a good way to use extra parameters",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]

		config, _ := ioutil.ReadFile("./.grid/config.json")
		_config, err := sjson.Set(string(config), "branchName", branchName)
		iferr(err)
		_config, err = sjson.Set(_config, "ifSync", true)
		iferr(err)
		writeFile("./.grid/config.json", _config)

	},
}
