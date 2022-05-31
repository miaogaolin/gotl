package response

import (
	"encoding/json"
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

func WriteSuccess(w http.ResponseWriter, data interface{})  {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(Response{Data: data})
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func WriteError(w http.ResponseWriter, err error)  {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(Response{Error: err.Error()})
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}