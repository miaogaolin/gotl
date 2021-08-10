package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogaolin/printlove-go/common/yaml2go"
	"github.com/miaogaolin/printlove-go/response"
)

func YamlToGo(c *gin.Context) {
	schema := c.PostForm("schema")
	yaml := yaml2go.New()
	data, err := yaml.Convert([]byte(schema))
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, data)
}
