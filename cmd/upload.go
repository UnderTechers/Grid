package cmd

import (
	"fmt"
	"grid/sha1_encode"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
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

type changesInFile struct {
	hashcode string
	_type    string
}

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

func compare2Files(file1 string, file2 string) bool { // true means different and vice versa
	return false
}

func createFile(filename string) {

	// del file
	os.Remove(filename)
	newFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
}

func createDir(filename string) {

	newpath := filename
	err := os.MkdirAll(newpath, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}

func writeFile(filename string, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0666)
	if err != nil {
		log.Print(err)
	}
}

var add = &cobra.Command{
	Use:   "add [the filename that will be submitted]",
	Short: "Add takes charge of upload files into one submit.",
	Long:  "Add takes charge of upload files into one submit.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filename = args[0]

		if filename == "." {

		} else {
			var originalFilePath = filename
			dataJson, err := ioutil.ReadFile("/.grid/config.json")
			iferr(err)
			latest := gjson.Get(string(dataJson), "latest").String()
			ifSync := gjson.Get(string(dataJson), "ifsync").Bool()
			username := gjson.Get(string(dataJson), "username").String()
			branchName := gjson.Get(string(dataJson), "banrchName").String()
			tmpConfigPath := "./.grid/tmp/config.json"
			// initialize tmp
			if ifSync == false { // if this is a new submit
				sjson.Set(string(dataJson), "ifsync", "true")
				res, _ := PathExists("./.grid/tmp/")
				if res { // if exists
					os.RemoveAll("./.grid/tmp/")
					createDir("./.grid/tmp/")
				}

				createFile(tmpConfigPath)
				writeFile(tmpConfigPath, "{}")
			}

			var targetedFilePath = path.Join(".", ".grid", username, branchName, latest, "file", originalFilePath)
			var d Diff
			_type := "normal"
			changes := make(map[string][]int)
			status := "changed"
			Config, err := ioutil.ReadFile(tmpConfigPath)
			commitConfig := string(Config)
			if res, _ := PathExists(targetedFilePath); !res {
				status = "add"
			}
			if d.If_Diff_Files(originalFilePath, targetedFilePath) {
				fileInfo := gjson.Get(commitConfig, originalFilePath).String() //json result

				hashcode := sha1_encode.ShaFile(originalFilePath)

				changes["addition"] = append(changes["addition"])
				// check if it exists
				if fileInfo == "" {
					content := make(map[string]string)
					content["hashcode"] = hashcode
					content["type"] = _type
					content["status"] = status
					sjson.Set(commitConfig, originalFilePath, content)
				}
			}
			fmt.Println(latest)

			//write back
			writeFile(tmpConfigPath, commitConfig)
			//compare
		}

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
		fmt.Println(title, content, dataJson)
		if ifSync == false {
			// error : because you cannot submit multiple times
			fmt.Println("- Error [101]: you cannot submit multiple times. Please check you have synced before you submit. ")
			return
		}
		fmt.Println(latest)

	},
}
