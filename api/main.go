package main

import (
	"fmt"
	"github.com/banerwai/gather/api/handler"
	"github.com/banerwai/gather/api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	router.SetupBaseRoute(r)

	router.SetupV1UserRoute(r)

	// Get user value
	r.GET("/user/:email", handler.GetUserByEmail)

	// Listen and Server in 0.0.0.0:3000
	fmt.Println("GIN Listening and serving HTTP on :3000")
	r.Run(":3000")
}
