package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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

var (
	hostPrefix = "127.0.0.1:8080/internal"
	client     = &http.Client{ //config to client by http
		Timeout: time.Second * 5,
	}
)

func iferr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getLatestBranchStatus() (string, bool, string) {
	dataJson, err := ioutil.ReadFile("/.grid/config.json")
	iferr(err)
	latest := gjson.Get(string(dataJson), "latest").String()
	ifSync := gjson.Get(string(dataJson), "ifsync").Bool()

	return latest, ifSync, string(dataJson)
}

var add = &cobra.Command{
	Use:   "add [the filename that will be submitted]",
	Short: "Add takes charge of upload files into one submit.",
	Long:  "Add takes charge of upload files into one submit.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filepath = args[0]
		latest, ifSync, dataJson := getLatestBranchStatus()
		if ifSync == false { // if this is a new submit
			dataJson, _ = sjson.Set(dataJson, "ifsync", "true")
			//create submit folder and corresponding json

		}
		fmt.Println(filepath)
	},
}

var submit = &cobra.Command{
	Use:   "submit [submit title] [submit content]",
	Short: "Submitting your new changes into a specific branch/version",
	Long:  "Submit command is used to update your changes in a specific branch/version. At the mean time, it means that sync command will be set as uploading mode",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var title = args[0]
		var content = args[1]

		latest, ifSync, dataJson := getLatestBranchStatus()

		if ifSync == false {
			// error : because you cannot submit multiple times
			fmt.Println("- Error [101]: you cannot submit multiple times. Please check you have synced before you submit. ")
			return
		}
		fmt.Println(latest)

	},
}
