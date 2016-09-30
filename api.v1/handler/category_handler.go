package handler

import (
	"github.com/banerwai/gather/common/flagparse"
	"github.com/banerwai/gather/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GET /category/:path?sign=xxx&timestamp=xxx
func LoadCategory(c *gin.Context) {
	_path := c.Params.ByName("path")

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _path, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.CategoryService
	_bool := _service.LoadCategory(_path)
	if !_bool {
		c.JSON(http.StatusOK, gin.H{"error": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": true})
}

// GET /category?sign=xxx&timestamp=xxx
func GetCategoriesBean(c *gin.Context) {

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.CategoryService
	_cats := _service.GetCategoriesBean()

	c.JSON(http.StatusOK, _cats)
}

// GET /category/:serialnumber?sign=xxx&timestamp=xxx
func GetSubCategoriesBean(c *gin.Context) {
	_serialnumber := c.Params.ByName("sn")
	if len(_serialnumber) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "serialnumber nil or len == 0"})
		return
	}

	if flagparse.BanerwaiAPICheckSign {
		_sign := c.Query("sign")
		_timestamp := c.Query("timestamp")

		if !ApiV1CheckSign(_sign, _serialnumber, _timestamp) {
			c.JSON(http.StatusOK, gin.H{"error": "sign error"})
			return
		}
	}

	var _service service.CategoryService
	_sn, err := strconv.Atoi(_serialnumber)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "serialnumber error"})
		return
	}
	_subcats := _service.GetSubCategoriesBean(int32(_sn))

	if len(_subcats) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "null"})
		return
	}

	c.JSON(http.StatusOK, _subcats)
}
