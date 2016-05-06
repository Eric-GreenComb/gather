package service

import (
	"testing"
)

func TestAuth(t *testing.T) {

	var _service AuthService
	s := _service.Login("ministor@126.com", "a11111")
	if s != "true" {
		t.Errorf("Login error")
	}

}
