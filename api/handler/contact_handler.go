package handler

import (
	"github.com/banerwai/gather/flagparse"
	"github.com/banerwai/gather/service"
	"github.com/banerwai/global/bean"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"
	"net/http"
)

// GET /contact/tpl/:tplname?sign=xxx&timestamp=xxx
func GetContactTpl(c *gin.Context) {
	_tplname := c.Params.ByName("tplname")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _tplname, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_tpl := _service.GetContactTpl(_tplname)

	c.JSON(http.StatusOK, _tpl)
}

// GET /contact/:id?sign=xxx&timestamp=xxx
func GetContactBean(c *gin.Context) {
	_id := c.Params.ByName("id")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_obj, _err := _service.GetContactBean(_id)
	if _err != nil {
		c.JSON(http.StatusOK, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusOK, _obj)
}

// GET /contact/:id/status?sign=xxx&timestamp=xxx
func GetContactSignStatusEnum(c *gin.Context) {
	_id := c.Params.ByName("id")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_obj := _service.GetContactSignStatusEnum(_id)

	c.JSON(http.StatusOK, _obj)
}

// GET /contact/client/:id?sign=xxx&timestamp=xxx
func GetClientContactBeans(c *gin.Context) {
	_email := c.Params.ByName("email")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _email, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_objs, _err := _service.GetClientContactBeans(_email)
	if _err != nil {
		c.JSON(http.StatusOK, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusOK, _objs)
}

// GET /contact/freelancer/:id?sign=xxx&timestamp=xxx
func GetFreelancerContactBeans(c *gin.Context) {
	_email := c.Params.ByName("email")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _email, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_objs, _err := _service.GetFreelancerContactBeans(_email)
	if _err != nil {
		c.JSON(http.StatusOK, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusOK, _objs)
}

// POST contact/add?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// client_email=client_email&freelancer_email=freelancer_email&contact_tpl=contact_tpl&tpl_param=tpl_param
func CreateContactBean(c *gin.Context) {

	_client_email := c.PostForm("client_email")
	_freelancer_email := c.PostForm("freelancer_email")
	_contact_tpl := c.PostForm("contact_tpl")
	_tpl_param := c.PostForm("tpl_param")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")
		if !ApiV1CheckSign(_sign, _client_email, _freelancer_email, _contact_tpl, _tpl_param, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _obj bean.Contact
	_obj.Id = bson.NewObjectId()
	_obj.ClientEmail = _client_email
	_obj.FreeLancerEmail = _freelancer_email
	_obj.ContactTpl = _contact_tpl
	_obj.TplParam = _tpl_param

	var _service service.ContactService
	_ret := _service.CreateContactBean(_obj)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST contact/update?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&param=param
func UpdateContact(c *gin.Context) {
	_id := c.PostForm("id")
	_param := c.PostForm("param")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _param, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_mmap_update := make(map[string]string)
	_mmap_update["tpl_param"] = _param
	_ret := _service.UpdateContact(_id, _mmap_update)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST contact/sign/client?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&status=status
func ClientSignContact(c *gin.Context) {
	_id := c.PostForm("id")
	_status := c.PostForm("status")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _status, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	bStatus := true
	if _status != "true" {
		bStatus = false
	}
	_ret := _service.ClientSignContact(_id, bStatus)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST contact/sign/freelance?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&status=status
func FreelancerSignContact(c *gin.Context) {
	_id := c.PostForm("id")
	_status := c.PostForm("status")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _status, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	bStatus := true
	if _status != "true" {
		bStatus = false
	}
	_ret := _service.FreelancerSignContact(_id, bStatus)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST contact/deal?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&status=true
func DealContact(c *gin.Context) {
	_id := c.PostForm("id")

	if flagparse.BanerwaiApiCheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_ret := _service.DealContact(_id, true)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}
