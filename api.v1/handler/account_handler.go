package handler

import (
	"github.com/banerwai/gather/common/flagparse"
	"github.com/banerwai/gather/service"
	"github.com/banerwai/global/bean"
	apiRequest "github.com/banerwai/global/bean"
	"github.com/banerwai/gommon/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GET /account/:uid
func GetAccountBean(c *gin.Context) {
	_uid := c.Params.ByName("uid")

	var _service service.AccountService
	_obj, _err := _service.GetAccountBean(_uid)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _obj)
}

// GET /billing/:id
func GetBillingBean(c *gin.Context) {
	_id := c.Params.ByName("id")

	var _service service.AccountService
	_obj, _err := _service.GetBillingBean(_id)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _obj)
}

// GET /billings/deal?uid=uid&stamp=stamp&pagesize=10
// stamp : page timestamp,for mongo page;
// timestamp : api timestamp
func GetDealBillingBeans(c *gin.Context) {
	_uid := c.Query("uid")
	_stamp_str := c.Query("stamp")
	_pagesize_str := c.Query("pagesize")

	_stamp, _err := strconv.ParseInt(_stamp_str, 10, 64)
	if _err != nil {
		_stamp = time.Now().Unix()
	}
	_pagesize, _err := strconv.ParseInt(_pagesize_str, 10, 64)
	if _err != nil {
		_pagesize = 10
	}

	var _service service.AccountService
	_objs, _err := _service.GetDealBillingBeans(_uid, _stamp, _pagesize)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _objs)
}

// GET /billings/all?uid=uid&stamp=stamp&pagesize=10
// stamp : page timestamp,for mongo page;
// timestamp : api timestamp
func GetBillingBeans(c *gin.Context) {
	_uid := c.Query("uid")
	_stamp_str := c.Query("stamp")
	_pagesize_str := c.Query("pagesize")

	_stamp, _err := strconv.ParseInt(_stamp_str, 10, 64)
	if _err != nil {
		_stamp = time.Now().Unix()
	}
	_pagesize, _err := strconv.ParseInt(_pagesize_str, 10, 64)
	if _err != nil {
		_pagesize = 10
	}

	var _service service.AccountService
	_objs, _err := _service.GetBillingBeans(_uid, _stamp, _pagesize)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _objs)
}

// CreateAccountBean create account handler
// POST account/add?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// user_id=user_id&email=email
func CreateAccountBean(c *gin.Context) {
	var apiReq apiRequest.APIRequest
	if c.BindJSON(&apiReq) != nil {
		c.JSON(http.StatusOK, gin.H{"error": "bind error"})
		return
	}

	if !api.CheckAPISignature(apiReq) {
		c.JSON(http.StatusOK, gin.H{"error": "sign error"})
		return
	}

	var _account bean.Account
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&_account) != nil {
		c.JSON(http.StatusOK, gin.H{"error": "bind data error"})
	}

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")
		if !ApiV1CheckSign(_sign, _account.UserID.Hex(), _account.Email, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.AccountService
	_ret := _service.CreateAccountBean(_account)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST /account/gen/:uid?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
func GenAccount(c *gin.Context) {
	_uid := c.Params.ByName("uid")

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _uid, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.AccountService
	_ret := _service.GenAccount(_uid)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST billing/add?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
// user_id=user_id&email=email
func CreateBillingBean(c *gin.Context) {
	var _billing bean.Billing
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&_billing) != nil {
		c.JSON(http.StatusOK, gin.H{"error": "bind data error"})
	}

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")
		if !ApiV1CheckSign(_sign, _billing.UserID.Hex(), _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.AccountService
	_ret := _service.CreateBillingBean(_billing)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST billing/deal/:bid?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
func DealBilling(c *gin.Context) {
	_bid := c.Params.ByName("bid")
	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _bid, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.AccountService
	_ret := _service.DealBilling(_bid)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

// POST billing/cancel/:bid?sign=xxx&timestamp=xxx HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
func CancelBilling(c *gin.Context) {
	_bid := c.Params.ByName("bid")

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _bid, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.AccountService
	_ret := _service.CancelBilling(_bid)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}
