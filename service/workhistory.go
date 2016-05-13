package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

type WorkHistoryService struct {
}

var _command_service command.WorkHistoryService
var _query_service query.WorkHistoryService

/**
 * command section
 */

func (self *WorkHistoryService) AddWorkHistory(json_workhistory string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.AddWorkHistory(json_workhistory)
	return
}

func (self *WorkHistoryService) UpdateWorkHistory(json_workhistory string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateWorkHistory(json_workhistory)
	return
}

/**
 * query section
 */
func (self *WorkHistoryService) GetWorkHistory(profile_id string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetWorkHistory(profile_id)
	return
}
