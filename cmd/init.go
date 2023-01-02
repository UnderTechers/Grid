package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tidwall/sjson"
)

var (
	newFile *os.File
	err     error
)

var Init = &cobra.Command{ // for testing cobra
	Use:   "init",
	Short: "Init is for creating a project in local grid server",
	Long: `Init will make a ".grid" folder in the root directory of project.
	 Every operations and changes will be uploaded to .grid folder first, add to submit and sync with server.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateDir("./.grid")
		CreateDir("./.grid/tmp")
		createFile("./.grid/config.json")
		createFile("./.grid/tmp/package.json")

		var source, branchName string
		fmt.Scanf("Input the source of this project(default=http://grid.gridle.com): %s", source)
		fmt.Scanf("Input the branchName(default=branch-dev): %s", branchName)
		config := "{}"
		sjson.Set(config, "source", source)
		sjson.Set(config, "branch-name", branchName)
		writeFile("./.grid/config.json", "{}")
		fmt.Println("- Initialization finished!")
		fmt.Println("- Please finish your configuration in ./config.json")

	},
}
