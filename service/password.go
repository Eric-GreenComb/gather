package api

import (
	"fmt"

	"github.com/banerwai/gather/service"
)

// 找回密码
// 修改密码
var overHours = 2.0 // 小时后过期

type PasswordService struct {
}

// 1. 找回密码, 通过email找用户,
// 用户存在, 生成code
func (self *PasswordService) FindPwd(email string) bool {

	fmt.Println(self.getRender())
	fmt.Println(self.getToken(email))

	return true
}

func (self *PasswordService) getRender() string {
	var _render service.RenderService
	_render_service, _ := _render.DefaultService()
	defer _render.CloseService()

	_r := _render_service.RenderHello("hello", "eric")

	return _r
}

func (self *PasswordService) getToken(email string) int64 {
	var _token service.TokenService
	_token_service, _ := _token.DefaultService()
	defer _token.CloseService()

	_verify := _token_service.VerifyToken(email, 1)

	return _verify
}
