package service

import (
	"github.com/banerwai/global/bean"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"

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
	_service, _ := _user_command_service.Default()
	defer _user_command_service.Close()
	v = _service.CreateUser(mmap)
	return
}

func (self *UserService) CreateBeanUser(user bean.User) (v string) {
	_mmap := make(map[string]string)
	_mmap["invited"] = user.Invited.String()
	_mmap["email"] = user.Email
	_mmap["pwd"] = user.Pwd
	return self.CreateUser(_mmap)
}

func (self *UserService) ResetPwd(email string, newpwd string) (v bool) {
	_service, _ := _user_command_service.Default()
	defer _user_command_service.Close()
	v = _service.ResetPwd(email, newpwd)
	return
}

func (self *UserService) ActiveUser(email string) (v bool) {
	_service, _ := _user_command_service.Default()
	defer _user_command_service.Close()
	v = _service.ActiveUser(email)
	return
}

/**
 * query section
 */

func (self *UserService) GetUser(email string) (v string) {
	_service, _ := _user_query_service.Default()
	defer _user_query_service.Close()
	v = _service.GetUser(email)
	return
}

func (self *UserService) GetDtoUser(email string) bean.UserDto {
	var _user bean.UserDto
	_data := self.GetUser(email)
	if len(_data) == 0 {
		return _user
	}
	bson.Unmarshal([]byte(_data), &_user)

	return _user
}

func (self *UserService) CountUser() (v int64) {
	_service, _ := _user_query_service.Default()
	defer _user_query_service.Close()
	v = _service.CountUser()
	return
}

/**
 * common section
 */
