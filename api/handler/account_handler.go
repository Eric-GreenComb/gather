package handler

import (
	"github.com/banerwai/gather/service"
	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strconv"
	"time"
)

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

func GetDealBillingBeans(c *gin.Context) {
	_uid := c.Query("uid")
	_timestamp_str := c.Query("timestamp")
	_pagesize_str := c.Query("pagesize")

	_timestamp, _err := strconv.ParseInt(_timestamp_str, 10, 64)
	if _err != nil {
		_timestamp := time.Now().Unix()
	}
	_pagesize, _err := strconv.ParseInt(_timestamp_str, 10, 64)
	if _err != nil {
		_pagesize = 10
	}

	var _service service.AccountService
	_objs, _err := _service.GetDealBillingBeans(_uid, _timestamp, _pagesize)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _objs)
}

func GetBillingBeans(c *gin.Context) {
	_uid := c.Query("uid")
	_timestamp_str := c.Query("timestamp")
	_pagesize_str := c.Query("pagesize")

	_timestamp, _err := strconv.ParseInt(_timestamp_str, 10, 64)
	if _err != nil {
		_timestamp := time.Now().Unix()
	}
	_pagesize, _err := strconv.ParseInt(_timestamp_str, 10, 64)
	if _err != nil {
		_pagesize = 10
	}

	var _service service.AccountService
	_objs, _err := _service.GetBillingBeans(_uid, _timestamp, _pagesize)
	if _err != nil {
		c.String(http.StatusOK, _err.Error())
		return
	}
	c.JSON(http.StatusOK, _objs)
}

func CreateAccountBean(c *gin.Context) {

	var _account bean.Account
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&_account) != nil {
		c.JSON(http.StatusOK, gin.H{"error": "bind data error"})
	}

	var _service service.AccountService
	_ret := _service.CreateAccountBean(_account)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

func GenAccount(c *gin.Context) {
	_uid := c.Params.ByName("uid")

	var _service service.AccountService
	_ret := _service.GenAccount(_uid)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

func CreateBillingBean(c *gin.Context) {

	var _billing bean.Billing
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&_billing) != nil {
		c.JSON(http.StatusOK, gin.H{"error": "bind data error"})
	}

	var _service service.AccountService
	_ret := _service.CreateBillingBean(_billing)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

func DealBilling(c *gin.Context) {
	_bid := c.Params.ByName("bid")

	var _service service.AccountService
	_ret := _service.DealBilling(_bid)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}

func CancelBilling(c *gin.Context) {
	_bid := c.Params.ByName("bid")

	var _service service.AccountService
	_ret := _service.CancelBilling(_bid)
	c.JSON(http.StatusOK, gin.H{"success": _ret})
}
