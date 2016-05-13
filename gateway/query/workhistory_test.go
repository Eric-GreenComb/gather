package query

import (
	"testing"
)

func TestWorkHistoryDefaultService(t *testing.T) {

	var _workhistory_service WorkHistoryService
	_thrift_service, _ := _workhistory_service.Default()

	v := _thrift_service.GetWorkHistory("1234567")

	if v != "1234567" {
		t.Errorf("GetWorkHistory error")
	}
	_workhistory_service.Close()

}
