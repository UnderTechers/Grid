package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

func DefaultBytes(maxBytes int64, description ...string) *progressbar.ProgressBar {
	desc := ""
	if len(description) > 0 {
		desc = description[0]
	}
	return progressbar.NewOptions64(
		maxBytes,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionUseANSICodes(true),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
	)
}

func downloadExample() {
	golangPkg := "https://golang.google.cn/dl/go1.16.4.darwin-amd64.pkg"
	req, _ := http.NewRequest("GET", golangPkg, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile("go1.16.4.darwin-amd64.pkg", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	var bar *progressbar.ProgressBar
	bar = progressbar.NewOptions(100, progressbar.OptionUseANSICodes(true))
	bar = DefaultBytes(
		resp.ContentLength,
		"- ",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)

}

var _sync = &cobra.Command{
	Use:   "sync",
	Short: "sync will make your submits synchronize with the server",

	Long: "If you want to synchrnoize the server because of new changes there, you can use sync. If you want to submit your changes, you can use sync.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("- sync starts!")
		fmt.Println("- detecting conflicts...")
		dataJson, err := ioutil.ReadFile("/.grid/config.json")
		iferr(err)
		ifsync := gjson.Get(string(dataJson), "ifsync").Bool()
		if ifsync {
			// upload
			// auth first
			projName := gjson.Get(string(dataJson), "projectName").String()
			userName := gjson.Get(string(dataJson), "username").String()
			token := gjson.Get(string(dataJson), "token").String()
			resp, err := client.PostForm(GridlePrefix+"/internal/auth", url.Values{"projName": {projName}, "username": {userName}, "token": {token}})
			defer resp.Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			if gjson.Get(string(body), "res").String() == "res" {
				// successful
				fmt.Println("- Auth succeeded!")
				// start to sync
			} else {
				fmt.Println("- Error[103] There is an error occurring when authenticating from server.")
				return
			}

			// upload the latest submit
			//

		} else {
			//download
			// the latest submit will be changed into the newest one
			// stagingArea will be outputed
		}
	},
}
