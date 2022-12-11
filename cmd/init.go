package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
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
		newpath := filepath.Join(".", ".grid")
		if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		fmt.Println("- Initialization finished!")
		fmt.Println("- Please finish your configuration in ./config.json")

	},
}
