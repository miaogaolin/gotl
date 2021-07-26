package main

import (
	"github.com/miaogaolin/printlove-go/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	handler.Route(r)
	Server(r, ":8080")
}
