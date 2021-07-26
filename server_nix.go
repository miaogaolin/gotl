// +build !windows

package main

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

func Server(r *gin.Engine, port string) error {
	return gracehttp.Serve(
		&http.Server{
			Addr:    ":" + port,
			Handler: r,
		},
	)
}
