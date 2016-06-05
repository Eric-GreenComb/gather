package service

import (
	"testing"

	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
)

func TestUser(t *testing.T) {

	var _service UserService

	// user_mmap := make(map[string]string)

	// user_mmap["email"] = "ministor@126.com"
	// user_mmap["pwd"] = "a11111"

	// v := _service.CreateUser(user_mmap)
	var _user bean.User
	_user.Invited = bson.ObjectIdHex("5707cb10ae6faa1d1071a189")
	_user.Email = "ministor@126.com"
	_user.Pwd = "a11111"

	v := _service.CreateBeanUser(_user)
	if !bson.IsObjectIdHex(v) {
		t.Errorf("CreateUser error")
	}

	b := _service.ActiveUser("ministor@126.com")
	if b != true {
		t.Errorf("ActiveUser error")
	}

	_dto, _ := _service.GetUserByEmailDto("ministor@126.com")
	if _dto.Email != "ministor@126.com" {
		t.Errorf("CountUser error")
	}

	n := _service.CountUser()
	if n != 100 {
		t.Errorf("CountUser error")
	}
}
