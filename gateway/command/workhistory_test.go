package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestWorkHistoryService(t *testing.T) {

	var _workhistory_service WorkHistoryService
	_thrift_service, _ := _workhistory_service.Default()

	v := _thrift_service.AddWorkHistory("email")

	if v != "email" {
		t.Errorf("AddWorkHistory error")
	}

	v = _thrift_service.UpdateWorkHistory("123456")

	if v != "123456" {
		t.Errorf("UpdateWorkHistory error")
	}
	_workhistory_service.Close()

}
