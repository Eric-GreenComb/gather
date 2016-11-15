package service

import (
	"errors"
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
)

// UserService UserService
type UserService struct {
}

var _userCommandService command.UserService
var _userQueryService query.UserService

/**
 * command section
 */

// CreateUser create user by map
func (us *UserService) CreateUser(mmap map[string]string) (v string) {
	_service, _err := _userCommandService.Default()
	if _err != nil {
		return
	}
	defer _userCommandService.Close()
	v = _service.CreateUser(mmap)
	return
}

// CreateBeanUser create user by user bean
func (us *UserService) CreateBeanUser(user bean.User) (v string) {
	_mmap := make(map[string]string)
	_mmap["invited"] = user.Invited.Hex()
	_mmap["email"] = user.Email
	_mmap["pwd"] = user.Pwd
	return us.CreateUser(_mmap)
}

// ResetPwd reset password
func (us *UserService) ResetPwd(email string, newpwd string) (v bool) {
	_service, _err := _userCommandService.Default()
	if _err != nil {
		return
	}
	defer _userCommandService.Close()
	v = _service.ResetPwd(email, newpwd)
	return
}

// ActiveUser active user
func (us *UserService) ActiveUser(email string) (v bool) {
	_service, _err := _userCommandService.Default()
	if _err != nil {
		return
	}
	defer _userCommandService.Close()
	v = _service.ActiveUser(email)
	return
}

/**
 * query section
 */

// GetUserByEmail get user json by email
func (us *UserService) GetUserByEmail(email string) (v string) {
	_service, _err := _userQueryService.Default()
	if _err != nil {
		return
	}
	defer _userQueryService.Close()
	v = _service.GetUserByEmail(email)
	return
}

// GetUserByEmailDto get user bean by email
func (us *UserService) GetUserByEmailDto(email string) (bean.UserDto, error) {
	var _user bean.UserDto
	_data := us.GetUserByEmail(email)
	if len(_data) == 0 {
		return _user, errors.New("get user by email is error")
	}
	bson.Unmarshal([]byte(_data), &_user)

	return _user, nil
}

// GetUserByID get user by ID
func (us *UserService) GetUserByID(ID string) (v string) {
	_service, _err := _userQueryService.Default()
	if _err != nil {
		return
	}
	defer _userQueryService.Close()
	v = _service.GetUserByID(ID)
	return
}

// GetUserByIDDto get user bean by ID
func (us *UserService) GetUserByIDDto(ID string) (bean.UserDto, error) {
	var _user bean.UserDto
	_data := us.GetUserByID(ID)
	if len(_data) == 0 {
		return _user, errors.New("get user by ID is error")
	}
	bson.Unmarshal([]byte(_data), &_user)

	return _user, nil
}

// CountUser count user
func (us *UserService) CountUser() (v int64) {
	_service, _err := _userQueryService.Default()
	if _err != nil {
		return
	}
	defer _userQueryService.Close()
	v = _service.CountUser()
	return
}

/**
 * common section
 */
