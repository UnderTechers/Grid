package cmd

import (
	"net/http"
	"time"
)

var (
	GridlePrefix = "127.0.0.1:8080/"
	client       = &http.Client{ //config to client by http
		Timeout: time.Second * 5,
	}
)
