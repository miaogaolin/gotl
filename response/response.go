package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Data: data})
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Response{Error: err.Error()})
}
