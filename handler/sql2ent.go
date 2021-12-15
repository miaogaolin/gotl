package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogaolin/gotl/response"
	sql2ent "github.com/miaogaolin/sql2ent/parser"
)

func SqlToEnt(c *gin.Context) {
	ddl := c.PostForm("ddl")
	res, err := sql2ent.Parse(ddl)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, res)
}
