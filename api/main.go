package main

import (
	"fmt"
	"github.com/banerwai/gather/api/router"
	"github.com/banerwai/gather/flagparse"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.SetupBaseRoute(r)

	router.SetupV1BussnessRoute(r)

	// Listen and Server in 0.0.0.0:3000
	fmt.Println("GIN Listening and serving HTTP on " + flagparse.BanerwaiApiPort)
	r.Run(flagparse.BanerwaiApiPort)
}
