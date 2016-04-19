package service

import (
	"testing"
)

// need start micro render service localhost:6020
func TestUserOpenService(t *testing.T) {

	var _user_service UserService
	_thrift_service, _ := _user_service.Default()

	v := _thrift_service.UpdatePwd("email", "pwd", "fromUserId")

	if v != true {
		t.Errorf("UpdatePwd error")
	}
	_user_service.Close()

	_user_service.Init()
	_thrift_service, _ = _user_service.Open()
	defer _user_service.Close()
	v2 := _thrift_service.CountUser()
	if v2 != 100 {
		t.Errorf("CountUser error")
	}
}
