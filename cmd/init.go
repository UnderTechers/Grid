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
		fmt.Printf("Input the source of this project(default=http://grid.gridle.com): ")
		fmt.Scanf("%s\n", &source)
		fmt.Printf("Input the branchName(default=branch-dev): ")
		fmt.Scanf("%s", &branchName)
		config := "{}"
		fmt.Println(source, branchName)
		var err error
		config, err = sjson.Set(config, "source", source)
		iferr(err)
		config, err = sjson.Set(config, "branchName", branchName)
		iferr(err)
		fmt.Println(config)
		writeFile("./.grid/config.json", config)
		fmt.Println("- Initialization finished!")
		fmt.Println("- Please finish your configuration in ./config.json")

	},
}
