package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "welcome 2 banerwai api")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
