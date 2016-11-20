package query

import (
	"fmt"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"testing"
)

func TestWorkHistoryDefaultService(t *testing.T) {

	var _workhistoryService WorkHistoryService
	_thriftService, _ := _workhistoryService.Default()

	v := _thriftService.GetWorkHistory("5707cb10ae6faa1d1071a189")

	var _workhistory bean.WorkHistory
	bson.Unmarshal([]byte(v), &_workhistory)
	fmt.Println(_workhistory)

	if len(_workhistory.HistoryAndFeedbacks) != 1 {
		t.Errorf("GetWorkHistory error")
	}
	_workhistoryService.Close()

}
