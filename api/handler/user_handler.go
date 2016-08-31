package handler

import (
	"github.com/banerwai/gather/flagparse"
	"github.com/banerwai/gather/service"
	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"
	"net/http"
)

// GET /user/:email?sign=xxx&timestamp=xxx
func GetUserByEmail(c *gin.Context) {
	_email := c.Params.ByName("email")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _email, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.UserService
	_user, _err := _service.GetUserByEmailDto(_email)
	if _err != nil {
		c.JSON(http.StatusOK, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusOK, _user)
}

// GET /user/:id?sign=xxx&timestamp=xxx
func GetUserByID(c *gin.Context) {
	_id := c.Params.ByName("id")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.UserService
	_user, _err := _service.GetUserByIDDto(_id)
	if _err != nil {
		c.JSON(http.StatusOK, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusOK, _user)
}

// POST /user/add?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// email=email&pwd=pwd&invited=
func CreateBeanUser(c *gin.Context) {

	_email := c.PostForm("email")
	_pwd := c.PostForm("pwd")
	_invited := c.PostForm("invited")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")
		if !ApiV1CheckSign(_sign, _email, _pwd, _invited, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _user bean.User

	_user.Email = _email
	_user.Pwd = _pwd
	if len(_invited) == 0 || !bson.IsObjectIdHex(_invited) {
		_user.Invited = bson.ObjectIdHex(global.BANERWAI_INVITED_DEFAULT_USERID)
	}

	var _service service.UserService
	_ret := _service.CreateBeanUser(_user)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST user/reset?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// email=email&newpwd=pwd
func ResetPwd(c *gin.Context) {
	_email := c.PostForm("email")
	_newpwd := c.PostForm("newpwd")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _email, _newpwd, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.UserService
	_ret := _service.ResetPwd(_email, _newpwd)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST user/active?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// email=email
func ActiveUser(c *gin.Context) {
	_email := c.PostForm("email")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _email, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.UserService
	_ret := _service.ActiveUser(_email)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}
