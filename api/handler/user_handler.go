package handler

import (
	"github.com/banerwai/gather/service"
	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func GetUserByEmail(c *gin.Context) {
	_email := c.Params.ByName("email")
	var _service service.UserService
	_user, _err := _service.GetUserByEmailDto(_email)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _user)
}

func GetUserByID(c *gin.Context) {
	_id := c.Params.ByName("id")
	var _service service.UserService
	_user, _err := _service.GetUserByIDDto(_id)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _user)
}

func CreateBeanUser(c *gin.Context) {

	var _user bean.User
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&_user) != nil {
		c.JSON(http.StatusOK, gin.H{"error": "bind User error"})
	}

	if len(_user.Invited) == 0 {
		_user.Invited = bson.ObjectIdHex(global.BANERWAI_INVITED_DEFAULT_USERID)
	}

	var _service service.UserService
	_ret := _service.CreateBeanUser(_user)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

func ResetPwd(c *gin.Context) {
	_email := c.Query("email")
	_newpwd := c.Query("newpwd")

	var _service service.UserService
	_ret := _service.ResetPwd(_email, _newpwd)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

func ActiveUser(c *gin.Context) {
	_email := c.Query("email")

	var _service service.UserService
	_ret := _service.ActiveUser(_email)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}
