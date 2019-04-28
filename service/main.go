package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	// Creates a router without any middleware by default
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	UseRouters(r)

	r.Run(":8081")
}
