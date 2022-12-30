package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
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
	_type := c.PostForm("type")
	if _type == "upload" {
		userName := c.PostForm("username")
		token := c.PostForm("token")
		projName := c.PostForm("projName")

		// auth to server
		urlValues := url.Values{} //parameters
		urlValues.Add("username", userName)
		urlValues.Add("token", token)
		urlValues.Add("projName", projName)
		rep, err := client.PostForm(hostPrefix+"/auth", urlValues)
		defer rep.Body.Close()
		iferr(err)

		body, err := ioutil.ReadAll(rep.Body)
		iferr(err)
		content := string(body)
		fmt.Println(content)

	}

	if _type == "download" {

	}

}
func (s *Router) Init_Server() {
	s.r = gin.New()
	defer s.r.Run(s.Host)
	s.Routes()

}

func (s Router) Routes() {
	s.r.POST("/proj/:username/:projectname", s.Download)
	s.r.POST("/sync", func(c *gin.Context) {
		// username, token(for authorize)
		// upload the 7z of project and then unzip it here
		// creating a new submit here and then save everything
		// and then updating the config.json in .grid_server
		// after done, send back the response of being successful

	})
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
