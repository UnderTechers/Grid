package cmd

import (
	"fmt"
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
	u, err := url.Parse(link)
	iferr(err)
	dirNameSeq := strings.Split(u.Path, "/")
	dirName := dirNameSeq[len(dirNameSeq)-1]
	fmt.Println(dirName)

	req, _ := http.NewRequest("GET", link, nil)
	resp, _ := http.DefaultClient.Do(req)
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

}
