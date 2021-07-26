package handler

import (
	"github.com/miaogaolin/printlove-go/response"

	"github.com/miaogaolin/printlove-go/common/sql2gorm/parser"

	"github.com/gin-gonic/gin"
)

func SqlToGorm(c *gin.Context) {
	ddl := c.PostForm("ddl")
	res, err := parser.ParseSqlFormat(ddl)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, string(res))
}
