package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

type UserService struct {
}

var _command_service command.UserService
var _query_service query.UserService

/**
 * command section
 */

func (self *UserService) CreateUser(mmap map[string]string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.CreateUser(mmap)
	return
}

func (self *UserService) ResetPwd(email string, newpwd string) (v bool) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.ResetPwd(email, newpwd)
	return
}

func (self *UserService) ActiveUser(email string) (v bool) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.ActiveUser(email)
	return
}

/**
 * query section
 */

func (self *UserService) GetUser(email string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetUser(email)
	return
}

func (self *UserService) CountUser() (v int64) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.CountUser()
	return
}

/**
 * common section
 */
