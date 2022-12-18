package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// type submitHistory struct {
// 	projectName    string   `json:"projectName"`
// 	branchHashCode string   `json:"branchHashCode"`
// 	branchName     string   `json:"branchName"`
// 	submits        string   `json:"submits"`
// 	latest         string   `json:"latest"`
// 	ifsync         bool     `json:"msg"`
// 	submit_token   string   `json:"submit_token"`
// 	contributors   []string `json:"contributors"`
// 	source         string   `json:"source"`
// }

func iferr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkSync() bool { //check if upload needed
	dataJson, err := ioutil.ReadFile("/.grid/config.json")
	iferr(err)
	latest := gjson.Get(string(dataJson), "latest")
	fmt.Println(latest)
	return false
}

var add = &cobra.Command{
	Use:   "add [the filename that will be submitted]",
	Short: "Add takes charge of upload files into one submit.",
	Long:  "Add takes charge of upload files into one submit.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filepath = args[0]
		checkSync()
		fmt.Println(filepath)
	},
}
