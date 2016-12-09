package router

import (
	"github.com/banerwai/gather/api.v1/handler"
	"github.com/gin-gonic/gin"
)

// SetupBaseRoute base route
func SetupBaseRoute(g *gin.Engine) {
	r := g.Group("/")
	{
		r.GET("", handler.Index)
		r.GET("/ping", handler.Ping)
	}
}

// SetupV1BussnessRoute bussness route
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
