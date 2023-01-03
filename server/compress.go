package server

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type Compression struct {
}

func writeFile(filename string, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0666)
	if err != nil {
		log.Print(err)
	}
}

func (c Compression) Compress(path string, filename string) {
	cmd := exec.Command("7z", "a", "-r", filename, path)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func (c Compression) Decompress(filename string) {

	if runtime.GOOS == "windows" {
		rootpath, _ := os.Getwd()

		writeFile("test.cmd", ("7z x " + "\"" + rootpath + "\\" + filename + "\""))
		cmd := exec.Command("zip")
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
