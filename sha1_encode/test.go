package sha1_encode

import (
	"log"
	"os"
	"path/filepath"
)

func Testify() {
	newpath := filepath.Join(".", ".grid")
	err := os.MkdirAll(newpath, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}
