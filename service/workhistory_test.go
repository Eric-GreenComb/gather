package service

import (
	"github.com/banerwai/global/bean"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestWorkHistoryBase(t *testing.T) {

	var _service WorkHistoryService

	var _defaultObjectID = "5707cb10ae6faa1d1071a189"
	var _obj bean.WorkHistory
	_obj.Id = bson.ObjectIdHex(_defaultObjectID)
	_obj.ProfileID = bson.ObjectIdHex(_defaultObjectID)

	var lsWorkHistoryAndFeedbacks []bean.WorkHistoryAndFeedback

	var _WorkHistoryAndFeedback1 bean.WorkHistoryAndFeedback
	_WorkHistoryAndFeedback1.Title = "ceshi"
	_WorkHistoryAndFeedback1.WorkPeriod = "2016.01-2016.04"
	_WorkHistoryAndFeedback1.WorkHours = 40

	var lsWorkFeedbacks []bean.WorkFeedback

	var _WorkFeedback1 bean.WorkFeedback
	_WorkFeedback1.WorkRate = 4
	_WorkFeedback1.Feedback = "perfect..."
	lsWorkFeedbacks = append(lsWorkFeedbacks, _WorkFeedback1)

	var _WorkFeedback2 bean.WorkFeedback
	_WorkFeedback2.WorkRate = 4
	_WorkFeedback2.Feedback = "good job..."

	lsWorkFeedbacks = append(lsWorkFeedbacks, _WorkFeedback2)

	_WorkHistoryAndFeedback1.WorkFeedbacks = lsWorkFeedbacks

	lsWorkHistoryAndFeedbacks = append(lsWorkHistoryAndFeedbacks, _WorkHistoryAndFeedback1)

	_obj.HistoryAndFeedbacks = lsWorkHistoryAndFeedbacks

	v := _service.UpdateWorkHistoryBean(_defaultObjectId, _obj)

	if v != "OK" {
		t.Errorf("UpdateWorkHistoryBean error")
	}

	_bean := _service.GetWorkHistoryBean(_defaultObjectId)

	if _bean.HistoryAndFeedbacks[0].WorkFeedbacks[0].WorkRate != 4 {
		t.Errorf("GetWorkHistoryBean error")
	}
}
