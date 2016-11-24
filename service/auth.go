package service

import (
	"errors"
	"github.com/banerwai/gather/gateway/query"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/gommon/regexp"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

// AuthService AuthService
type AuthService struct {
}

var _authQueryService query.AuthService

/**
 * query section
 */

// Login Login return user json
func (as *AuthService) Login(email string, pwd string) (string, error) {
	if len(email) == 0 || len(pwd) == 0 {
		return "", errors.New("input is null")
	}
	if !regexp.IsEmail(email) {
		return "", errors.New("email regexp is error ")
	}
	_service, _err := _authQueryService.Default()
	if _err != nil {
		return "", _err
	}
	defer _authQueryService.Close()
	v := _service.Login(email, pwd)
	if strings.Contains(v, "error:") {
		return "", errors.New(v)
	}
	return v, nil
}

// LoginDto return user bean
func (as *AuthService) LoginDto(email string, pwd string) (bean.UserDto, error) {
	_v, _err := as.Login(email, pwd)
	var _user bean.UserDto
	if _err != nil {
		return _user, _err
	}
	if strings.Contains(_v, "error:") {
		return _user, errors.New(_v)
	}
	bson.Unmarshal([]byte(_v), &_user)
	return _user, nil
}

/**
 * common section
 */

// // 找回密码
// // 修改密码
// var overHours = 2.0 // 小时后过期

// type PasswordService struct {
// }

// // 1. 找回密码, 通过email找用户,
// // 用户存在, 生成code
// func (self *PasswordService) FindPwd(email string) bool {

// 	fmt.Println(self.getRender())
// 	fmt.Println(self.getToken(email))

// 	return true
// }

// func (self *PasswordService) getRender() string {
// 	var _render service.RenderService
// 	_render_service, _ := _render.DefaultService()
// 	defer _render.CloseService()

// 	_r := _render_service.RenderHello("hello", "eric")

// 	return _r
// }

// func (self *PasswordService) getToken(email string) int64 {
// 	var _token service.TokenService
// 	_token_service, _ := _token.DefaultService()
// 	defer _token.CloseService()

// 	_verify := _token_service.VerifyToken(email, 1)

// 	return _verify
// }
