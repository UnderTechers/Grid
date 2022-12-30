package main

import (
	"fmt"
	"grid/cmd"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {

	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// fmt.Printf("fileHash2 = %x \n", encrypt.ShaFile("./test.txt"))
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// encrypt.Testify()
	// fmt.Println(hex.EncodeToString(encrypt.ShaFile2()))

	// var c server.Compression
	// c.Compress("hello", "something.7z")
	//cmd.Execute()
	fmt.Println(cmd.GetFiles("./.grid/tmp"))

	// var r server.Router
	// r.Host = "0.0.0.0:8000"
	// r.Init_Server()
	// cmd.DownloadFile("https://grid.gridle.com/xcloudfance/ProjectName")
}
