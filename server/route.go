package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var (
	hostPrefix = "127.0.0.1:8080/internal"
	client     = &http.Client{ //config to client by http
		Timeout: time.Second * 5,
	}
)

func init() {

}

type Router struct {
	Host string
	r    *gin.Engine
}

func iferr(err error) { //grammar sugar for error detection
	if err != nil {
		log.Fatalln(err)
	}
}

func sync(c *gin.Context) { //POST
	//_type := c.PostForm("type")

	_file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Get form err: %s", err.Error()))
		return
	}
	basePath := "./upload/"
	filebase := filepath.Base(_file.Filename)
	filename := basePath + filebase
	filesuffix := path.Ext(filename)
	fileprefix := filebase[0 : len(filebase)-len(filesuffix)]

	if err := c.SaveUploadedFile(_file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("- Error[106] upload file err: %s", err.Error()))
		return
	}

	//unzip
	var cmpress Compression
	cmpress.Decompress(filename)
	target := basePath + "/" + fileprefix
	session := c.Request.Header.Get("session")
	userName := c.Request.Header.Get("userName")
	projName := c.Request.Header.Get("projName")
	branchName := c.Request.Header.Get("branchName")
	//auth

	// auth to server
	urlValues := url.Values{} //parameters
	urlValues.Add("username", userName)
	urlValues.Add("token", session)
	urlValues.Add("projName", projName)
	rep, err := client.PostForm(hostPrefix+"/auth", urlValues)
	defer rep.Body.Close()
	iferr(err)
	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if gjson.Get(string(body), "res").String() == "successful" {
		// successful
		fmt.Println("- Auth succeeded!")
		// start to sync
		os.Rename(target, path.Join(".", ".grid_server", userName, projName, branchName, fileprefix))
		//update config.json

		_config, _ := ioutil.ReadFile("./.grid/config.json")
		config := string(_config)
		config, _ = sjson.Set(config, "latest", fileprefix)
		c.String(http.StatusOK, fmt.Sprintf("- Project synced successfully. ", _file.Filename))

	}

}
func (s *Router) Init_Server() {
	s.r = gin.New()
	defer s.r.Run(s.Host)
	s.Routes()

}

func (s Router) Routes() {
	s.r.POST("/:username/:projectname", s.Download)
	s.r.POST("/sync", sync)
	//func(c *gin.Context) {
	// username, token(for authorize)

	// upload the 7z of project and then unzip it here
	// creating a new submit here and then save everything
	// and then updating the config.json in .grid_server
	// after done, send back the response of being successfulrr76666e6e6e6e55

	//}

}

func (s Router) Download(c *gin.Context) {
	fmt.Println("succeeded")
	userName := c.Param("username")
	projectName := c.Param("projectname")
	dir := "./.grid_server/" + userName + "/" + projectName + "/"
	// -- get default branch //

	f, _ := ioutil.ReadFile(dir + "config.json")
	_default := gjson.Get(string(f), "defaultBranch").String()

	//latest submit
	g, _ := ioutil.ReadFile(dir + _default + "/history.json")
	_latest := gjson.Get(string(g), "latest").String()

	fmt.Println(dir + _default + "/" + _latest + ".7z")
	// compress
	var compress Compression
	compress.Compress(dir+_default+"/"+_latest+"/files/", dir+_default+"/"+_latest+".7z")
	// about compression: it would be moved when every submit updated

	// send file
	c.File(dir + _default + "/" + _latest + ".7z")

}
