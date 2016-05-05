package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestUserOpenService(t *testing.T) {

	var _user_service UserService
	_thrift_service, _ := _user_service.Default()

	v := _thrift_service.ResetPwd("ministor@126.com", "pwd")

	if v != true {
		t.Errorf("ResetPwd error")
	}
	_user_service.Close()

	_user_service.Init()
	_thrift_service, _ = _user_service.Open()
	defer _user_service.Close()
	v2 := _thrift_service.ActiveUser("ministor@126.com")
	if v2 != true {
		t.Errorf("ActiveUser error")
	}
}
