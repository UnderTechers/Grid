package cmd

import "github.com/spf13/cobra"

var download = &cobra.Command{
	Use:   "download [link to download]",
	Short: "Download the project you havent got before",

	Long: "Download command is used to download the project that users haven't had it. Also after downloading into a specific folder, the project is initialized locally",
}
