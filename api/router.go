package main

import (
	"github.com/banerwai/gather/api/handler"
	"github.com/gin-gonic/gin"
)

func setupBaseRoute(g *gin.Engine) {
	r := g.Group("/")
	{
		r.GET("", handler.Index)
		r.GET("/ping", handler.Ping)
	}
}

func setupV1Route(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("", handler.V1)
		r.GET("/user/:email", handler.GetUserByEmail)
	}
}
