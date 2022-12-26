package cmd

import (
	"grid/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listen)
	listen.Flags().StringVarP(&port, "port", "p", "8000", "Port set to be listened")
}

var port string

var listen = &cobra.Command{
	Use:   "listen [Host]",
	Short: "Open HTTP listening to transform Grid into a tiny server",
	Long:  "Open HTTP listening to transform Grid into a tiny server. It should be taken 1 parameter to set port(seeing the -h documentation)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var r server.Router
		r.Host = args[0] + ":" + port
		r.Init_Server()

	},
}
