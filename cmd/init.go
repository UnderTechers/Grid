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

func init_tmp() {
	os.RemoveAll("./.grid/tmp")
	CreateDir("./.grid/tmp")
	createFile("./.grid/tmp/package.json")
	createFile("./.grid/tmp/stagingArea.json")
	writeFile("./.grid/tmp/stagingArea.json", "{}")
}

var Init = &cobra.Command{ // for testing cobra
	Use:   "init",
	Short: "Init is for creating a project in local grid server",
	Long: `Init will make a ".grid" folder in the root directory of project.
	 Every operations and changes will be uploaded to .grid folder first, add to submit and sync with server.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateDir("./.grid")
		createFile("./.grid/config.json")
		init_tmp()

		var source, branchName string
		fmt.Scanf("Input the source of this project(default=http://grid.gridle.com): %s", source)
		fmt.Scanf("Input the branchName(default=branch-dev): %s", branchName)
		config := "{}"
		sjson.Set(config, "source", source)
		sjson.Set(config, "branch-name", branchName)
		writeFile("./.grid/config.json", config)
		fmt.Println("- Initialization finished!")
		fmt.Println("- Please finish your configuration in ./config.json")

	},
}
