package handler

import (
	"bytes"

	"github.com/basgys/goxml2json"
	"github.com/gin-gonic/gin"
	"github.com/miaogaolin/printlove-go/response"
)

func XmlToJson(c *gin.Context) {
	schema := c.PostForm("schema")
	content := bytes.NewReader([]byte(schema))
	b, err := xml2json.Convert(content,
		xml2json.WithTypeConverter(xml2json.Float))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, b.String())

}
