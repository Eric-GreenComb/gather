package service

import (
	"testing"
)

// need start micro render service localhost:6020
func TestTokenOpenService(t *testing.T) {

	var _auth_service AuthService
	_thrift_service, _ := _auth_service.Default()

	v := _thrift_service.Register("email", "pwd", "fromUserId")

	if v != "emailpwdfromUserId" {
		t.Errorf("Register error")
	}
	_auth_service.Close()

	_auth_service.Init()
	_thrift_service, _ = _auth_service.Open()
	defer _auth_service.Close()
	v2 := _thrift_service.Login("emailOrUsername", "pwd")
	if v2 != "emailOrUsernamepwd" {
		t.Errorf("Login error")
	}
}
