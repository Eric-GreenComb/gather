package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestUserOpenService(t *testing.T) {

	var _userService UserService
	_thriftService, _ := _userService.Default()

	v := _thriftService.ResetPwd("ministor@126.com", "pwd")

	if v != true {
		t.Errorf("ResetPwd error")
	}
	_userService.Close()

	_userService.Init()
	_thriftService, _ = _userService.Open()
	defer _userService.Close()
	v2 := _thriftService.ActiveUser("ministor@126.com")
	if v2 != true {
		t.Errorf("ActiveUser error")
	}
}
