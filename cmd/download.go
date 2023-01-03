package cmd

import (
	"fmt"
	"grid/server"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var download = &cobra.Command{
	Use:   "download [link to download]",
	Short: "Download the project you havent got before",

	Long: "Download command is used to download the project that users haven't had it. Also after downloading into a specific folder, the project is initialized locally",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		DownloadFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(download)
}
func DownloadFile(link string) {
	if strings.HasPrefix(link, "http://") != true && strings.HasPrefix(link, "https://") != true {
		link = "http://" + link
	}
	u, err := url.Parse(link)
	iferr(err)

	// source := u.Host
	// _dataJson, _ := ioutil.ReadFile("./.grid/config.json")
	// dataJson, _ := sjson.Set(string(_dataJson), "source", source)
	// writeFile("./.grid/config.json", dataJson)

	dirNameSeq := strings.Split(u.Path, "/")
	dirName := dirNameSeq[len(dirNameSeq)-1]
	fmt.Println(dirName)

	req, err := http.NewRequest("POST", link, nil)
	resp, _ := http.DefaultClient.Do(req)
	iferr(err)
	defer resp.Body.Close()

	f, _ := os.OpenFile(dirName+".7z", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	var bar *progressbar.ProgressBar
	bar = progressbar.NewOptions(100, progressbar.OptionUseANSICodes(true))
	bar = DefaultBytes(
		resp.ContentLength,
		"- ",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)

	var c server.Compression

	c.Decompress("./" + dirName + ".7z")

	err = os.Rename("./files", "./"+dirName)
	iferr(err)

}
