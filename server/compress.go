package server

import (
	"log"
	"os"
	"os/exec"
)

type Compression struct {
}

func (c Compression) Compress(path string, filename string) {
	cmd := exec.Command("7z", "a", "-r", filename, path)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func (c Compression) Decompress(filename string, target string) {
	os.RemoveAll(target)
	cmd := exec.Command("7z", "x", filename, "-o"+target)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
