package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"

	"encoding/json"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
)

type WorkHistoryService struct {
}

var _command_service command.WorkHistoryService
var _query_service query.WorkHistoryService

/**
 * command section
 */

func (self *WorkHistoryService) UpdateWorkHistory(prifile_id, workhistory string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateWorkHistory(prifile_id, workhistory)
	return
}

func (self *WorkHistoryService) UpdateWorkHistoryBean(prifile_id string, workhistory bean.WorkHistory) (v string) {
	b, err := json.Marshal(workhistory)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateWorkHistory(prifile_id, string(b))
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

func (self *WorkHistoryService) GetWorkHistoryBean(profile_id string) (v bean.WorkHistory) {
	_workhistory := self.GetWorkHistory(profile_id)
	bson.Unmarshal([]byte(_workhistory), &v)
	return
}
