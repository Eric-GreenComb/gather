package service

import (
	"errors"
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
)

type UserService struct {
}

var _user_command_service command.UserService
var _user_query_service query.UserService

/**
 * command section
 */

func (self *UserService) CreateUser(mmap map[string]string) (v string) {
	_service, _err := _user_command_service.Default()
	if _err != nil {
		return
	}
	defer _user_command_service.Close()
	v = _service.CreateUser(mmap)
	return
}

func (self *UserService) CreateBeanUser(user bean.User) (v string) {
	_mmap := make(map[string]string)
	_mmap["invited"] = user.Invited.Hex()
	_mmap["email"] = user.Email
	_mmap["pwd"] = user.Pwd
	return self.CreateUser(_mmap)
}

func (self *UserService) ResetPwd(email string, newpwd string) (v bool) {
	_service, _err := _user_command_service.Default()
	if _err != nil {
		return
	}
	defer _user_command_service.Close()
	v = _service.ResetPwd(email, newpwd)
	return
}

func (self *UserService) ActiveUser(email string) (v bool) {
	_service, _err := _user_command_service.Default()
	if _err != nil {
		return
	}
	defer _user_command_service.Close()
	v = _service.ActiveUser(email)
	return
}

/**
 * query section
 */

func (self *UserService) GetUserByEmail(email string) (v string) {
	_service, _err := _user_query_service.Default()
	if _err != nil {
		return
	}
	defer _user_query_service.Close()
	v = _service.GetUserByEmail(email)
	return
}

func (self *UserService) GetUserByEmailDto(email string) (bean.UserDto, error) {
	var _user bean.UserDto
	_data := self.GetUserByEmail(email)
	if len(_data) == 0 {
		return _user, errors.New("get user by email is error")
	}
	bson.Unmarshal([]byte(_data), &_user)

	return _user, nil
}

func (self *UserService) GetUserByID(id string) (v string) {
	_service, _err := _user_query_service.Default()
	if _err != nil {
		return
	}
	defer _user_query_service.Close()
	v = _service.GetUserByID(id)
	return
}

func (self *UserService) GetUserByIDDto(id string) (bean.UserDto, error) {
	var _user bean.UserDto
	_data := self.GetUserByID(id)
	if len(_data) == 0 {
		return _user, errors.New("get user by id is error")
	}
	bson.Unmarshal([]byte(_data), &_user)

	return _user, nil
}

func (self *UserService) CountUser() (v int64) {
	_service, _err := _user_query_service.Default()
	if _err != nil {
		return
	}
	defer _user_query_service.Close()
	v = _service.CountUser()
	return
}

/**
 * common section
 */
