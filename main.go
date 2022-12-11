package main

import (
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

	cmd.Execute()

}
