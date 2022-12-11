package server

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	host string
}

func (s Router) Init_Server() {
	r := gin.New()

	defer r.Run(s.host)

}
