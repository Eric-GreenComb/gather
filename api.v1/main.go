package main

import (
	"fmt"
	"github.com/banerwai/gather/api.v1/router"
	"github.com/banerwai/gather/common/flagparse"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.SetupBaseRoute(r)

	router.SetupV1BussnessRoute(r)

	// Listen and Server in 0.0.0.0:3000
	fmt.Println("GIN Listening and serving HTTP on " + flagparse.BanerwaiAPIPort)
	r.Run(flagparse.BanerwaiAPIPort)
}
