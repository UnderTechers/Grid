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

func init() {
	rootCmd.AddCommand(add)
	rootCmd.AddCommand(rm)
}

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

func GetFiles(folder string, extraDir string) []string {

	files, _ := ioutil.ReadDir(folder + extraDir)
	var ret []string
	for _, file := range files {
		if file.IsDir() {
			ret = append(ret, GetFiles(folder, "/"+extraDir+"/"+file.Name())...)
		} else {
			ret = append(ret, path.Clean((extraDir + "/" + file.Name())))
		}
	}
	return ret

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

var rm = &cobra.Command{
	Use:   "rm [the filename you want to remove from the submit]",
	Short: "Remove the files from this submit",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// do it later
	},
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
			tmpConfigPath := "./.grid/tmp/stagingArea.json"
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

			var targetedFilePath = path.Join(".", ".grid", username, branchName, latest, "files", originalFilePath)
			var d Diff
			_type := "normal"
			hashcode := sha1_encode.ShaFile(originalFilePath)
			changes := make(map[string][]int)
			status := "changed"
			Config, err := ioutil.ReadFile(tmpConfigPath)
			commitConfig := string(Config)
			if res, _ := PathExists(targetedFilePath); !res {
				status = "add"
				content := make(map[string]string)
				content["hashcode"] = hashcode
				content["type"] = _type
				content["status"] = status
				sjson.Set(commitConfig, originalFilePath, content)
				writeFile(tmpConfigPath, commitConfig)
				return
			}
			if d.If_Diff_Files(originalFilePath, targetedFilePath) {
				//check if exists
				value := gjson.Get(commitConfig, originalFilePath)
				if !value.Exists() {
					//error
					fmt.Println("- Error[102] : There is no file in given directory")
					return
				}
				fileInfo := value.String() //json result

				changes["addition"] = append(changes["addition"])
				// check if it exists
				if fileInfo == "" {
					content := make(map[string]string)
					content["hashcode"] = hashcode
					content["type"] = _type
					content["status"] = status
					sjson.Set(commitConfig, originalFilePath, content)
					writeFile(tmpConfigPath, commitConfig)
					return
				}
			}

			//write back
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

		data, err := ioutil.ReadFile("/.grid/config.json") //get info from config.json, which shows the default settings
		iferr(err)
		dataJson := string(data)
		defaultBranch := gjson.Get(dataJson, "branchName").String()
		username := gjson.Get(dataJson, "username").String()
		submit_token := gjson.Get(dataJson, "submit_token").String()
		latestSubmit := gjson.Get(dataJson, "latest").String()
		ifSync := gjson.Get(dataJson, "ifsync").Bool()

		fmt.Println(title, content, dataJson)
		if ifSync == false {
			// error : because you cannot submit multiple times
			fmt.Println("- Error [101]: you cannot submit multiple times. Please check you have synced before you submit. ")
			return
		} else {

			stagingArea, err := ioutil.ReadFile("./.grid/tmp/stagingArea.json") // read changes
			if err != nil {
				log.Fatalln(err)
			}
			fileList := GetFiles("./.grid/tmp", "")
			addition := 0
			deletion := 0
			modification := 0

			for _, v := range fileList {
				// iterate all of changes in stagingArea
				if gjson.Get(string(stagingArea), v).Exists() {
					// if exists
					status := gjson.Get(string(stagingArea), v).Get("status").String()
					if status == "add" {
						addition += 1
					}
					if status == "changed" {
						modification += 1
					}
					if status == "deleted" {
						deletion += 1
					}
				}
			}

			fmt.Println("(%d) files added, (%d) files changed, (%d) files deleted", addition, modification, deletion)
			//create a submit

			submitHashCode := sha1_encode.ShaText(string(stagingArea))
			targetedPath := path.Join(".", ".grid", defaultBranch, submitHashCode)
			createDir(targetedPath)
			createDir(targetedPath + "/files")

			//update latest submit
			sjson.Set(dataJson, "latest", submitHashCode)
			writeFile("./.grid/config.json", dataJson)
		}

	},
}
