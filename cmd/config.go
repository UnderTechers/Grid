package cmd

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/tidwall/sjson"
)

var username, token string

func init() {
	config.Flags().StringVarP(&token, "token", "t", "", "the token logined by Gridle")
	config.Flags().StringVarP(&username, "username", "u", "", "the username logined by Gridle")
	rootCmd.AddCommand(config)

}

var config = &cobra.Command{
	Use:   "config",
	Short: "to configure your tokens and username",
	Run: func(cmd *cobra.Command, args []string) {
		if username != "" {
			_dataJson, _ := ioutil.ReadFile("./.grid/config.json")
			dataJson, _ := sjson.Set(string(_dataJson), "username", username)
			writeFile("./.grid/config.json", dataJson)
		}
		if token != "" {
			_dataJson, _ := ioutil.ReadFile("./.grid/config.json")
			dataJson, _ := sjson.Set(string(_dataJson), "submit_token", token)
			writeFile("./.grid/config.json", dataJson)
		}
	},
}
