package handler

import (
	"github.com/banerwai/gather/common/flagparse"
	"github.com/banerwai/gather/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login POST /auth/login?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// email=email&pwd=pwd
func Login(c *gin.Context) {
	_email := c.PostForm("email")
	_pwd := c.PostForm("pwd")

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")
		if !ApiV1CheckSign(_sign, _email, _pwd, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.AuthService
	_obj, _err := _service.LoginDto(_email, _pwd)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _obj)
}
