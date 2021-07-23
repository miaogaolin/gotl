package handler

import (
	"printlove-go/response"

	"printlove-go/common/sql2gorm/parser"

	"github.com/gin-gonic/gin"
)

func SqlToGorm(c *gin.Context) {
	ddl := c.PostForm("ddl")
	res, err := parser.ParseSql(ddl)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, res)
}
