package command

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"testing"
)

// need start micro render service localhost:36020
func TestWorkHistoryService(t *testing.T) {

	var _workhistory_service WorkHistoryService
	_thrift_service, _ := _workhistory_service.Default()

	var _defaultObjectId = "5707cb10ae6faa1d1071a189"
	var _obj bean.WorkHistory
	_obj.Id = bson.ObjectIdHex(_defaultObjectId)
	_obj.ProfileID = bson.ObjectIdHex(_defaultObjectId)

	var lsWorkHistoryAndFeedbacks []bean.WorkHistoryAndFeedback

	var _WorkHistoryAndFeedback1 bean.WorkHistoryAndFeedback
	_WorkHistoryAndFeedback1.Title = "ceshi"
	_WorkHistoryAndFeedback1.WorkPeriod = "2016.01-2016.04"
	_WorkHistoryAndFeedback1.WorkHours = 40

	var lsWorkFeedbacks []bean.WorkFeedback

	var _WorkFeedback1 bean.WorkFeedback
	_WorkFeedback1.WorkRate = 3
	_WorkFeedback1.Feedback = "perfect"
	lsWorkFeedbacks = append(lsWorkFeedbacks, _WorkFeedback1)

	var _WorkFeedback2 bean.WorkFeedback
	_WorkFeedback2.WorkRate = 3
	_WorkFeedback2.Feedback = "good job"

	lsWorkFeedbacks = append(lsWorkFeedbacks, _WorkFeedback2)

	_WorkHistoryAndFeedback1.WorkFeedbacks = lsWorkFeedbacks

	lsWorkHistoryAndFeedbacks = append(lsWorkHistoryAndFeedbacks, _WorkHistoryAndFeedback1)

	_obj.HistoryAndFeedbacks = lsWorkHistoryAndFeedbacks

	b, _ := json.Marshal(_obj)

	v := _thrift_service.UpdateWorkHistory(_defaultObjectId, string(b))

	if v != "OK" {
		t.Errorf("UpdateWorkHistory error")
	}
	_workhistory_service.Close()

}
