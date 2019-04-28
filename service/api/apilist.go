package api


import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiList 显示所有路由列表
func ApiList(eng *gin.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		routes := eng.Routes()
		total := len(routes)
		urls := map[string]string{}
		for i := 0; i < total; i++ {
			r := routes[i]
			v, exists := urls[r.Path]
			if exists {
				v = v + ", " + r.Method
			} else {
				v = r.Method
			}
			urls[r.Path] = v
		}
		ctx.JSON(http.StatusOK, urls)
	}
}

