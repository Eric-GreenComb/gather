package handler

import (
	"github.com/banerwai/gather/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	_email := c.Query("email")
	_pwd := c.Query("pwd")
	var _service service.AuthService
	_obj, _err := _service.LoginDto(_email, _pwd)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _obj)
}
