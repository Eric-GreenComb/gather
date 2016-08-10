package handler

import (
	"github.com/banerwai/gather/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByEmail(c *gin.Context) {
	_email := c.Params.ByName("email")
	var _service service.UserService
	_user, _ := _service.GetUserByEmailDto(_email)
	c.JSON(http.StatusOK, _user)
}
