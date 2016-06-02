package service

import (
	"testing"
)

func TestAuth(t *testing.T) {

	var _service AuthService
	_user, _err := _service.LoginDto("ministor@126.com", "a11111")
	if _err != nil {
		t.Errorf(_err.Error())
	}
	if _user.Email != "ministor@126.com" {
		t.Errorf("user dto error")
	}
}
