package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// V1 version
func V1(c *gin.Context) {
	c.String(http.StatusOK, "api v1")
}
