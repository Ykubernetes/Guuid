package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ykubernetes/GUUID/Service"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.GET("/get", Service.CreateHandler)
	r.GET("/get/simple", Service.CreateSimpleHandler)
	r.GET("/mget/:num", Service.CreateMultiHandler)
	r.GET("/mget/:num/simple", Service.CreateMultiSimpleHandler)
	return r
}
