package main

import (
	"xinxin/service/api"
	"github.com/gin-gonic/gin"
)

func UseRouters(eng *gin.Engine) {
	rg := eng.Group("/")

	rg.GET("/", api.ApiList(eng))
}
