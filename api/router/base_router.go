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

func SetupV1BussnessRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("", handler.V1)
	}

	SetupV1UserRoute(g)
	SetupV1AccountRoute(g)
	SetupV1AuthRoute(g)

	SetupV1CategoryRoute(g)

	SetupV1ContactRoute(g)
}
