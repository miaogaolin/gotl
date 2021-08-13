package handler

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {

	r.POST("/sql2gorm", SqlToGorm)
	r.POST("/sql2gozero", SqlToGoZero)
	r.POST("sql2ent", SqlToEnt)
	r.POST("yaml2go", YamlToGo)
	r.POST("xml2json", XmlToJson)

}
