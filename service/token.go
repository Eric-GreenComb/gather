package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
	"github.com/banerwai/global/bean"
)

// TokenService TokenService
type TokenService struct {
}

var _tokenCommandService command.TokenService
var _tokenQueryService query.TokenService

/**
 * command section
 */

// CreateToken create a new token
func (ts *TokenService) CreateToken(key string, ttype int64) (v string) {
	_service, _err := _tokenCommandService.Default()
	if _err != nil {
		return _err.Error()
	}
	defer _tokenCommandService.Close()
	v = _service.CreateToken(key, ttype)
	return
}

// DeleteToken delete token
func (ts *TokenService) DeleteToken(key string, ttype int64) (v bool) {
	_service, _err := _tokenCommandService.Default()
	if _err != nil {
		return
	}
	defer _tokenCommandService.Close()
	v = _service.DeleteToken(key, ttype)
	return
}

/**
 * query section
 */

// VerifyToken verify token
// return -1 不存在
// return -2 过期
// return -3 db error
// return 1 验证pass
func (ts *TokenService) VerifyToken(token string, ttype int64, overhour int64) (v int64) {
	_service, _err := _tokenQueryService.Default()
	if _err != nil {
		return
	}
	defer _tokenQueryService.Close()
	v = _service.VerifyToken(token, ttype, overhour)
	return
}

/**
 * common section
 */

// GetOverHours get overtime hours
func (ts *TokenService) GetOverHours(ttype int64) int64 {
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
