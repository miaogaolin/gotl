// +build windows

package main

import (
	"github.com/gin-gonic/gin"
)

func Server(r *gin.Engine, port string) error {
	return r.Run(":" + port)
}
