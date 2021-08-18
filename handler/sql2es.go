package handler

import (
	"github.com/cch123/elasticsql"
	"github.com/gin-gonic/gin"
	"github.com/miaogaolin/printlove-go/response"
)

func SqlToEs(c *gin.Context) {
	schema := c.PostForm("schema")
	res, _, err := elasticsql.Convert(schema)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, res)
}
