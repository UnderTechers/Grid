package cmd

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"grid/sha1_encode"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var diffCommand = &cobra.Command{
	Use:   "diff [file1] [file2]",
	Short: "To compare 2 different files by their sha-1 code or using extensions in Grid",
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

func GetFiles(folder string) []string {
	files, _ := ioutil.ReadDir(folder)
	var result []string
	for _, file := range files {
		if file.IsDir() {
			result = append(result)

		} else {
			fmt.Println(folder + "/" + file.Name())
		}
	}
	return result
}

func (d Diff) If_Diff_Files(filepath1 string, filepath2 string) bool {
	sha1_code := hex.EncodeToString(sha1_encode.ShaFile(filepath1))
	sha2_code := hex.EncodeToString(sha1_encode.ShaFile(filepath2))
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
	preCode, err := os.Open(preCodePath)
	iferr(err)
	postCode, err := os.Open(postCodePath)
	iferr(err)
	defer preCode.Close()
	defer postCode.Close()

	scanner1 := bufio.NewScanner(preCode)
	scanner2 := bufio.NewScanner(postCode)

	for scanner1.Scan() && scanner2.Scan() {
		line1 := strip(scanner1.Text())
		line2 := strip(scanner2.Text())
		if line1 == "" && line2 != "" {
			// new
		}
		if line1 != "" && line2 == "" {
			//delete
		}
		if line1 == line2 {
			continue
		}

		if line1 != line2 {
			//changed
		}

	}

}
