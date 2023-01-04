package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func init() {

}

// 以键值对形式上传文件
func postFile2(uri, filePath, requestHost, session, userName, projName, branchName string) {

	paramName := "file"

	//打开要上传的文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	defer file.Close()
	body := &bytes.Buffer{}
	//创建一个multipart类型的写文件
	writer := multipart.NewWriter(body)
	//使用给出的属性名paramName和文件名filePath创建一个新的form-data头
	part, err := writer.CreateFormFile(paramName, filePath)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	//将源复制到目标，将file写入到part   是按默认的缓冲区32k循环操作的，不会将内容一次性全写入内存中,这样就能解决大文件的问题
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		fmt.Println(" post err=", err)
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("S-COOKIE2", "a=2l=310260000000000&m=460&n=00")
	request.Header.Add("session", session)
	request.Header.Add("userName", userName)
	request.Header.Add("projName", projName)
	request.Header.Add("branchName", branchName)
	//writer.FormDataContentType() ： 返回w对应的HTTP multipart请求的Content-Type的值，多以multipart/form-data起始
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//设置host，只能用request.Host = “”，不能用request.Header.Add(),也不能用request.Header.Set()来添加host
	request.Host = requestHost
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	clt := http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
	defer clt.CloseIdleConnections()
	res, err := clt.Do(request)
	//http返回的response的body必须close,否则就会有内存泄露
	defer func() {
		res.Body.Close()
		fmt.Println("- upload finished")
	}()
	if err != nil {
		fmt.Println("- Error[106] Error is ", err)
	}
	body1, err1 := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  err is ", err1)
		return
	}
	fmt.Println(string(body1[:]))
}
