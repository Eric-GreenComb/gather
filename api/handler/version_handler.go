package handler

import (
	"github.com/banerwai/gather/flagparse"
	"github.com/banerwai/gommon/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func V1(c *gin.Context) {
	c.String(http.StatusOK, "api v1")
}

func ApiV1CheckSign(sign string, args ...string) bool {
	if len(sign) == 0 {
		return false
	}
	total := flagparse.BanerwaiApiKey
	for _, arg := range args {
		total += arg
	}
	return crypto.CompareDoubleMd5(total, sign)
}
