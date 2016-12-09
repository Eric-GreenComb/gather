package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index index
func Index(c *gin.Context) {
	c.String(http.StatusOK, "welcome 2 banerwai api")
}

// Ping ping
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
