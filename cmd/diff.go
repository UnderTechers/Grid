package cmd

import (
	"fmt"
	"grid/sha1_encode"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var diffCommand = &cobra.Command{
	Use:   "diff [file1] [file2]",
	Short: "To compare 2 different files by their sha-1 code or using extensions in Grid",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var d Diff
		if !d.If_Diff_Files(args[0], args[1]) {
			fmt.Println("same")
		} else {
			fmt.Println("different")
		}

	},
}

func init() {
	rootCmd.AddCommand(diffCommand)
}

type Diff struct {
}

func PathExists(path string) (bool, error) { //used to determine whether file/path exists
	_, err := os.Stat(path)
	if err == nil { // if found, return true
		return true, nil
	}
	if os.IsNotExist(err) { //if not found, return false
		return false, nil
	}
	return false, err
}

func (d Diff) Get_latest_filePath(filePath string) string {
	var ret string
	dataJson, err := ioutil.ReadFile("/.grid/config.json")
	iferr(err)
	latest := gjson.Get(string(dataJson), "latest")
	ret = filepath.Join(".", ".grid", latest.String())

	var ifExist bool
	ifExist, err = PathExists(ret)
	iferr(err)
	if !ifExist {
		return ""
	}
	return ret
}

func (d Diff) If_Diff_Files(filepath1 string, filepath2 string) bool {
	sha1_code := sha1_encode.ShaFile(filepath1)
	sha2_code := (sha1_encode.ShaFile(filepath2))
	if sha1_code == sha2_code {
		return false
	} else {
		return true
	}
}

// 利用正则表达式压缩字符串，去除空格或制表符
func strip(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func (d Diff) Show_Changes(preCodePath string, postCodePath string) {
	// show the changes line by line
	cmd := exec.Command("npm", "diff", "--diff="+preCodePath, "--diff="+postCodePath)
	if err := cmd.Run(); err != nil {
		fmt.Println("- Error [400] Some unknown errors occurred.")
	}

}
