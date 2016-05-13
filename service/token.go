package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
	"github.com/banerwai/global/bean"
)

type TokenService struct {
}

var _command_service command.TokenService
var _query_service query.TokenService

/**
 * command section
 */

func (self *TokenService) NewToken_(key string, ttype int64) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.NewToken_(key, ttype)
	return
}

func (self *TokenService) DeleteToken(key string, ttype int64) (v bool) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.DeleteToken(key, ttype)
	return
}

/**
 * query section
 */

// return -1 不存在
// return -2 过期
// return -3 db error
// return 1 验证pass
func (self *TokenService) VerifyToken(token string, ttype int64, overhour float64) (v int64) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.VerifyToken(token, ttype, overhour)
	return
}

/**
 * common section
 */

func (self *TokenService) GetOverHours(ttype int64) float64 {
	switch ttype {

	// 0 - 2.0
	case bean.TokenPwd:
		return bean.PwdOverHours

	// 1 - 48.0
	case bean.TokenActiveEmail:
		return bean.ActiveEmailOverHours

	// 2 - 2.0
	case bean.TokenUpdateEmail:
		return bean.UpdateEmailOverHours
	}
	return 0
}
