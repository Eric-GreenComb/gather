package router

import (
	"github.com/banerwai/gather/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupBaseRoute(g *gin.Engine) {
	r := g.Group("/")
	{
		r.GET("", handler.Index)
		r.GET("/ping", handler.Ping)
	}
}
