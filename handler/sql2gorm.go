package handler

import (
	"github.com/miaogaolin/gotl/response"

	"github.com/miaogaolin/gotl/common/sql2gorm/parser"

	"github.com/gin-gonic/gin"
)

func SqlToGorm(c *gin.Context) {
	ddl := c.PostForm("ddl")
	res, err := parser.ParseSqlFormat(ddl,
		parser.WithGormType(),
		parser.WithJsonTag(),
	)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, string(res))
}
