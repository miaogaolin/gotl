package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogaolin/printlove-go/response"
	"github.com/miaogaolin/sql2ent"
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
