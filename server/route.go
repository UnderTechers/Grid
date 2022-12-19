package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	hostPrefix = "127.0.0.1:8080/internal"
	client     = &http.Client{ //config to client by http
		Timeout: time.Second * 5,
	}
)

type Router struct {
	Host string
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
func (s Router) Init_Server() {
	r := gin.New()

	defer r.Run(s.Host)

}
