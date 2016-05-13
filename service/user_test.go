package service

import (
	"testing"

	"github.com/banerwai/global/bean"
)

func TestUser(t *testing.T) {

	var _service UserService

	// user_mmap := make(map[string]string)

	// user_mmap["email"] = "ministor@126.com"
	// user_mmap["pwd"] = "a11111"

	// v := _service.CreateUser(user_mmap)
	var _user bean.User
	_user.Invited = "5707cb10ae6faa1d1071a189"
	_user.Email = "ministor@126.com"
	_user.Pwd = "a11111"

	v := _service.CreateBeanUser(_user)
	if v != "OK" {
		t.Errorf("CreateUser error")
	}

	b := _service.ActiveUser("ministor@126.com")
	if b != true {
		t.Errorf("ActiveUser error")
	}

	_dto := _service.GetDtoUser("ministor@126.com")
	if _dto.Email != "ministor@126.com" {
		t.Errorf("CountUser error")
	}

	n := _service.CountUser()
	if n != 100 {
		t.Errorf("CountUser error")
	}
}
