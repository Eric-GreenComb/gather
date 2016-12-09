package router

import (
	"github.com/banerwai/gather/api.v1/handler"
	"github.com/gin-gonic/gin"
)

// SetupV1AuthRoute auth route
func SetupV1AuthRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.POST("/auth/login", handler.Login)
	}
}
