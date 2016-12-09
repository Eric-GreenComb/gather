package handler

import (
	"github.com/banerwai/gather/common/flagparse"
	"github.com/banerwai/gather/service"
	"github.com/banerwai/global/bean"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"
	"net/http"
)

// GetContactTpl GET /contact/tpl/:tplname?sign=xxx&timestamp=xxx
func GetContactTpl(c *gin.Context) {
	_tplname := c.Params.ByName("tplname")

	if flagparse.BanerwaiAPICheckSign {
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

// GetContactBean GET /contact/:id?sign=xxx&timestamp=xxx
func GetContactBean(c *gin.Context) {
	_id := c.Params.ByName("id")

	if flagparse.BanerwaiAPICheckSign {
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

// GetContactSignStatusEnum GET /contact/:id/status?sign=xxx&timestamp=xxx
func GetContactSignStatusEnum(c *gin.Context) {
	_id := c.Params.ByName("id")

	if flagparse.BanerwaiAPICheckSign {
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

// GetClientContactBeans GET /contact/client/:id?sign=xxx&timestamp=xxx
func GetClientContactBeans(c *gin.Context) {
	_email := c.Params.ByName("email")

	if flagparse.BanerwaiAPICheckSign {
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

// GetFreelancerContactBeans GET /contact/freelancer/:id?sign=xxx&timestamp=xxx
func GetFreelancerContactBeans(c *gin.Context) {
	_email := c.Params.ByName("email")

	if flagparse.BanerwaiAPICheckSign {
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

// CreateContactBean POST contact/add?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// client_email=client_email&freelancer_email=freelancer_email&contact_tpl=contact_tpl&tpl_param=tpl_param
func CreateContactBean(c *gin.Context) {

	_clientEmail := c.PostForm("client_email")
	_freelancerEmail := c.PostForm("freelancer_email")
	_contactTpl := c.PostForm("contact_tpl")
	_tplParam := c.PostForm("tpl_param")

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")
		if !ApiV1CheckSign(_sign, _clientEmail, _freelancerEmail, _contactTpl, _tplParam, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _obj bean.Contact
	_obj.ID = bson.NewObjectId()
	_obj.ClientEmail = _clientEmail
	_obj.FreeLancerEmail = _freelancerEmail
	_obj.ContactTpl = _contactTpl
	_obj.TplParam = _tplParam

	var _service service.ContactService
	_ret := _service.CreateContactBean(_obj)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// UpdateContact POST contact/update?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&param=param
func UpdateContact(c *gin.Context) {
	_id := c.PostForm("id")
	_param := c.PostForm("param")

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _id, _param, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.ContactService
	_mmapUpdate := make(map[string]string)
	_mmapUpdate["tpl_param"] = _param
	_ret := _service.UpdateContact(_id, _mmapUpdate)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// ClientSignContact POST contact/sign/client?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&status=status
func ClientSignContact(c *gin.Context) {
	_id := c.PostForm("id")
	_status := c.PostForm("status")

	if flagparse.BanerwaiAPICheckSign {
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

// FreelancerSignContact POST contact/sign/freelance?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&status=status
func FreelancerSignContact(c *gin.Context) {
	_id := c.PostForm("id")
	_status := c.PostForm("status")

	if flagparse.BanerwaiAPICheckSign {
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

// DealContact POST contact/deal?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// id=id&status=true
func DealContact(c *gin.Context) {
	_id := c.PostForm("id")

	if flagparse.BanerwaiAPICheckSign {
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
