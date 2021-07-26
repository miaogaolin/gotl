package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogaolin/printlove-go/common/sql2gozero/gen"
	"github.com/miaogaolin/printlove-go/response"
	"github.com/tal-tech/go-zero/tools/goctl/config"
)

func SqlToGoZero(c *gin.Context) {
	ddl := c.PostForm("ddl")
	cache := c.PostForm("cache")
	g, err := gen.NewGenerator("model", &config.Config{NamingFormat: config.DefaultFormat})
	if err != nil {
		response.Error(c, err)
		return
	}

	isCache := false
	if cache == "1" {
		isCache = true
	}

	res, err := g.GenFromDDContent([]byte(ddl), isCache, "")
	if err != nil {
		response.Error(c, err)
		return
	}

	for _, v := range res {
		response.Success(c, v)
		return
	}
}
