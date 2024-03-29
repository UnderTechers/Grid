package cmd

import (
	"fmt"
	"grid/sha1_encode"
	"grid/utils"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func init() {
	rootCmd.AddCommand(add)
	rootCmd.AddCommand(submit)
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
	_dataJson, err := ioutil.ReadFile("/.grid/config.json")
	dataJson := string(_dataJson)
	iferr(err)
	latest := gjson.Get(dataJson, "latest").String()
	ifSync := gjson.Get(dataJson, "ifsync").Bool()

	return latest, ifSync, dataJson
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

func CreateDir(filename string) {
	os.RemoveAll(filename)
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
			_dataJson, err := ioutil.ReadFile("./.grid/config.json")
			iferr(err)
			dataJson := string(_dataJson)
			latest := gjson.Get(dataJson, "latest").String()

			ifSync := gjson.Get(dataJson, "ifsync").Bool()
			//username := gjson.Get(dataJson, "username").String()
			branchName := gjson.Get(dataJson, "branchName").String()
			tmpConfigPath := "./.grid/tmp/stagingArea.json"
			// initialize tmp
			if ifSync == false { // if this is a new submit
				dataJson, _ = sjson.Set(dataJson, "ifsync", true)
				defer writeFile("./.grid/config.json", dataJson)
			}

			var targetedFilePath = path.Join(".", ".grid", branchName, latest, "files", originalFilePath)
			var d Diff
			_type := "normal"
			hashcode := sha1_encode.ShaFile(originalFilePath)
			//changes := make(map[string][]int)

			Config, err := ioutil.ReadFile(tmpConfigPath)
			commitConfig := string(Config)

			fmt.Println(originalFilePath, targetedFilePath)

			if res, _ := PathExists(targetedFilePath); !res {
				status := "add"
				content := make(map[string]string)
				content["hashcode"] = hashcode
				content["type"] = _type
				content["status"] = status

				commitConfig, _ = sjson.Set(commitConfig, deal(originalFilePath), content)
				defer writeFile(tmpConfigPath, commitConfig)
				utils.Copy_Folder(originalFilePath, "./.grid/tmp/"+originalFilePath)
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

				//changes["addition"] = append(changes["addition"])
				// check if it exists
				if fileInfo == "" {
					status := "changed"
					content := make(map[string]string)
					content["hashcode"] = hashcode
					content["type"] = _type
					content["status"] = status
					commitConfig, _ = sjson.Set(commitConfig, deal(originalFilePath), content)
					defer writeFile(tmpConfigPath, commitConfig)
					utils.Copy_Folder(originalFilePath, "./.grid/tmp/"+originalFilePath)
					return
				}
			} else {
				fmt.Println("- Nothing different.")
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

		data, err := ioutil.ReadFile("./.grid/config.json") //get info from config.json, which shows the default settings
		iferr(err)
		dataJson := string(data)
		defaultBranch := gjson.Get(dataJson, "branchName").String()
		// username := gjson.Get(dataJson, "username").String()
		// submit_token := gjson.Get(dataJson, "submit_token").String()
		// latestSubmit := gjson.Get(dataJson, "latest").String()
		ifSync := gjson.Get(dataJson, "ifsync").Bool()

		//fmt.Println(title, content, dataJson)
		if ifSync == false {
			// error : because you cannot submit multiple times
			fmt.Println("- Error [101]: you cannot submit multiple times. Please check you have synced before you submit. ")
			return
		} else {
			dataJson, _ = sjson.Set(dataJson, "ifsync", false)
			writeFile("./.grid/config.json", dataJson)

			stagingArea, err := ioutil.ReadFile("./.grid/tmp/stagingArea.json") // read changes
			if err != nil {
				log.Fatalln(err)
			}
			fileList := GetFiles("./.grid/tmp", "")
			addition := 0
			deletion := 0
			modification := 0

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "Status", "Type"})

			for _, v := range fileList {
				// iterate all of changes in stagingArea
				base := deal(path.Base(v))
				//fmt.Println(string(stagingArea))
				fmt.Println(gjson.Get(string(stagingArea), base).Get("status"))
				if gjson.Get(string(stagingArea), base).Exists() {
					// if exists
					fmt.Println(1)
					status := gjson.Get(string(stagingArea), base).Get("status").String()
					_type := gjson.Get(string(stagingArea), base).Get("type").String()
					inline := []string{v, status, _type}
					table.Append(inline)

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

			fmt.Printf("- (%d) files added, (%d) files changed, (%d) files deleted\n", addition, modification, deletion)
			table.Render()
			//create a submit
			submitHashCode := sha1_encode.ShaText(string(stagingArea))
			targetedPath := path.Join(".", ".grid", defaultBranch, submitHashCode)

			//fmt.Println(defaultBranch, targetedPath)

			CreateDir(targetedPath)
			CreateDir(targetedPath + "/files")
			createFile(targetedPath + "/files/package.json") // for comparison by nodejs

			//update latest submit
			dataJson, _ = sjson.Set(dataJson, "latest", submitHashCode)
			writeFile("./.grid/config.json", dataJson)

			fmt.Println("- submit has been created! The SHA-1 code is: ", submitHashCode)

			//copy files
			utils.Copy_Folder("./", "../caches/")
			utils.Copy_Folder("./../caches/", targetedPath+"/files/")
			utils.Copy_Folder("./.grid/tmp", targetedPath+"/files/")

			utils.Cut(targetedPath+"/files/stagingArea.json", targetedPath+"/stagingArea.json") // cut it to become the stagingArea

			//update title and descriptions
			newStaging, _ := ioutil.ReadFile(targetedPath + "stagingArea.json")
			newStaging_, _ := sjson.Set(string(newStaging), "title", title)
			newStaging_, _ = sjson.Set(newStaging_, "desciption", content)
			writeFile(targetedPath+"stagingArea.json", newStaging_)

			//update latest
			_dataJson, _ := ioutil.ReadFile("./.grid/config.json")
			dataJson, _ := sjson.Set(string(_dataJson), "latest", submitHashCode)
			writeFile("./.grid/config.json", dataJson)
			fmt.Println("- submit conducted successfully")
			// clean tmp folder
			init_tmp()

		}

	},
}

var history = &cobra.Command{
	Use: "history",
	// to-do: to print the table of current submit of changes
}
