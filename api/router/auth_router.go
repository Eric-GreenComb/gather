package router

import (
	"github.com/banerwai/gather/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupV1AuthRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.POST("/auth/login", handler.Login)
	}
}
