package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

type ProfileService struct {
}

var _command_service command.ProfileService
var _query_service query.ProfileService

/**
 * command section
 */

func (self *ProfileService) AddProfile(json_profile string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.AddProfile(json_profile)
	return
}

func (self *ProfileService) UpdateProfile(json_profile string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateProfile(json_profile)
	return
}

func (self *ProfileService) DeleteProfile(id string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.DeleteProfile(id)
	return
}

/**
 * query section
 */

func (self *ProfileService) GetProfile(id string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetProfile(id)
	return
}

func (self *ProfileService) GetProfilesByEmail(email string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetProfilesByEmail(email)
	return
}

func (self *ProfileService) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	return
}
