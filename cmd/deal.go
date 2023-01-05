package cmd

import "strings"

func deal(dir string) string {
	return strings.ReplaceAll(dir, ".", "\\.")

}
