package main

import (
	"github.com/miaogaolin/printlove-go/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	handler.Route(r)
	r.Run(":8080")

}
