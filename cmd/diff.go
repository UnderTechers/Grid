package cmd

import (
	"encoding/hex"
	"grid/sha1_encode"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

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
	sha1_code := hex.EncodeToString(sha1_encode.ShaFile(filepath1))
	sha2_code := hex.EncodeToString(sha1_encode.ShaFile(filepath2))
	if sha1_code == sha2_code {
		return false
	} else {
		return true
	}
}
