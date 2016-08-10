package main

import (
	"github.com/banerwai/gather/api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	setupBaseRoute(r)
	setupV1Route(r)

	// Get user value
	r.GET("/user/:email", handler.GetUserByEmail)

	// Listen and Server in 0.0.0.0:3000
	r.Run(":3000")
}
