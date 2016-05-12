package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

type ResumeService struct {
}

var _command_service command.ResumeService
var _query_service query.ResumeService

/**
 * command section
 */

func (self *ResumeService) AddResume(json_resume string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.AddResume(json_resume)
	return
}

func (self *ResumeService) UpdateResume(json_resume string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateResume(json_resume)
	return
}

/**
 * query section
 */
func (self *ResumeService) GetResume(id string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetResume(id)
	return
}
