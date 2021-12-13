package main

import (
	"flag"

	"github.com/miaogaolin/printlove-go/handler"

	"github.com/gin-gonic/gin"
)

var post = flag.String("p", "8080", "port")

func main() {
	flag.Parse()
	r := gin.New()
	handler.Route(r)
	Server(r, *post)
}
