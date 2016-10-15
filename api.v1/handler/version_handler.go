package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func V1(c *gin.Context) {
	c.String(http.StatusOK, "api v1")
}
