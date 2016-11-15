package service

import (
	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"

	"encoding/json"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
)

// WorkHistoryService WorkHistoryService
type WorkHistoryService struct {
}

var _workhistoryCommandService command.WorkHistoryService
var _workhistoryQueryService query.WorkHistoryService

/**
 * command section
 */

// UpdateWorkHistory update user work history by profileID
func (whs *WorkHistoryService) UpdateWorkHistory(profileID, workhistory string) (v string) {
	_service, _err := _workhistoryCommandService.Default()
	if _err != nil {
		return
	}
	defer _workhistoryCommandService.Close()
	v = _service.UpdateWorkHistory(profileID, workhistory)
	return
}

// UpdateWorkHistoryBean update user work history beans by profileID
func (whs *WorkHistoryService) UpdateWorkHistoryBean(profileID string, workhistory bean.WorkHistory) (v string) {
	b, err := json.Marshal(workhistory)
	if err != nil {
		return err.Error()
	}
	v = whs.UpdateWorkHistory(profileID, string(b))
	return
}

/**
 * query section
 */

//GetWorkHistory get work history json by profileID
func (whs *WorkHistoryService) GetWorkHistory(profileID string) (v string) {
	_service, _err := _workhistoryQueryService.Default()
	if _err != nil {
		return
	}
	defer _workhistoryQueryService.Close()
	v = _service.GetWorkHistory(profileID)
	return
}

// GetWorkHistoryBean get work history bean by profileID
func (whs *WorkHistoryService) GetWorkHistoryBean(profileID string) (v bean.WorkHistory) {
	_workhistory := whs.GetWorkHistory(profileID)
	bson.Unmarshal([]byte(_workhistory), &v)
	return
}
