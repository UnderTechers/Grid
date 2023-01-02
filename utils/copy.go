package utils

import (
	"bufio"
	"fmt"
	"log"
	"path"

	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateDir(filename string) {
	os.RemoveAll(filename)
	newpath := filename
	err := os.MkdirAll(newpath, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}

func Copy(src, dst string, BUFFERSIZE int64) error {
	os.RemoveAll(src)
	os.RemoveAll(dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

func Copy_Folder(from, to string) error {

	f, e := os.Stat(from)

	if e != nil {
		return e
	}

	//to-do : ignore command
	if path.Base("to") == ".gridConfig" || path.Base("to") == "package.json" { //ignore
		return nil
	}
	if f.IsDir() {

		//from是文件夹，那么定义to也是文件夹
		CreateDir(to)
		if list, e := ioutil.ReadDir(from); e == nil {

			for _, item := range list {

				if e = Copy_Folder(filepath.Join(from, item.Name()), filepath.Join(to, item.Name())); e != nil {
					return e

				}

			}

		}

	} else {

		//from是文件，那么创建to的文件夹

		p := filepath.Dir(to)

		if _, e = os.Stat(p); e != nil {

			if e = os.MkdirAll(p, 0777); e != nil {
				return e
			}

		}
		//读取源文件

		file, e := os.Open(from)

		if e != nil {
			return e
		}

		defer file.Close()

		bufReader := bufio.NewReader(file)

		// 创建一个文件用于保存
		out, e := os.Create(to)
		if e != nil {
			return e
		}
		defer out.Close()
		// 然后将文件流和文件流对接起来
		_, e = io.Copy(out, bufReader)
	}
	return e

}

func Cut(from, to string) error {

	f, e := os.Stat(from)

	if e != nil {
		return e
	}

	//to-do : ignore command
	if path.Base("to") == ".gridConfig" || path.Base("to") == "package.json" { //ignore
		return nil
	}
	if f.IsDir() {

		//from是文件夹，那么定义to也是文件夹
		CreateDir(to)
		if list, e := ioutil.ReadDir(from); e == nil {

			for _, item := range list {

				if e = Copy_Folder(filepath.Join(from, item.Name()), filepath.Join(to, item.Name())); e != nil {
					return e

				}

			}

		}

	} else {

		//from是文件，那么创建to的文件夹

		p := filepath.Dir(to)

		if _, e = os.Stat(p); e != nil {

			if e = os.MkdirAll(p, 0777); e != nil {
				return e
			}

		}
		//读取源文件

		file, e := os.Open(from)

		if e != nil {
			return e
		}

		defer file.Close()

		bufReader := bufio.NewReader(file)

		// 创建一个文件用于保存
		out, e := os.Create(to)
		if e != nil {
			return e
		}
		defer out.Close()
		// 然后将文件流和文件流对接起来
		_, e = io.Copy(out, bufReader)
	}

	os.Remove(from)
	return e

}
