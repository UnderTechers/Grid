package cmd

import (
	"fmt"
	"grid/sha1_encode"

	"github.com/spf13/cobra"
)

var hash = &cobra.Command{
	Use:   "hash [filename]",
	Short: "get the SHA-1 code of the file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(sha1_encode.ShaFile(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(hash)
}
