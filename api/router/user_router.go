package router

import (
	"github.com/banerwai/gather/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupV1UserRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("", handler.V1)
		r.GET("/user/:email", handler.GetUserByEmail)
	}
}
