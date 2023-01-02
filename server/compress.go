package server

import (
	"log"
	"os/exec"
	"runtime"
)

type Compression struct {
}

func (c Compression) Compress(path string, filename string) {
	cmd := exec.Command("7z", "a", "-r", filename, path)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func (c Compression) Decompress(filename string, done func()) {
	defer done()

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd.exe", "/c start 7z"+" "+"x"+" "+filename)
		if err := cmd.Run(); err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		cmd := exec.Command("7z", "x", filename)
		if err := cmd.Run(); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
