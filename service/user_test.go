package service

import (
	"testing"
)

func TestUser(t *testing.T) {

	var _service UserService

	user_mmap := make(map[string]string)

	user_mmap["email"] = "ministor@126.com"
	user_mmap["pwd"] = "a11111"

	v := _service.CreateUser(user_mmap)
	if v != "OK" {
		t.Errorf("CreateUser error")
	}

	b := _service.ActiveUser("ministor@126.com")
	if b != true {
		t.Errorf("ActiveUser error")
	}

	n := _service.CountUser()
	if n != 100 {
		t.Errorf("CountUser error")
	}
}
