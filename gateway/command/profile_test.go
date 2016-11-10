package command

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"testing"
)

// need start micro render service localhost:36020
func TestProfileService(t *testing.T) {

	var _profileService ProfileService
	_thriftService, _ := _profileService.Default()

	_defaultObjectID := "5707cb10ae6faa1d1071a189"

	var _obj bean.Profile
	_obj.ID = bson.ObjectIdHex(_defaultObjectID)
	_obj.UserID = bson.ObjectIdHex(_defaultObjectID)
	_obj.Name = "Test"
	_obj.JobTitle = "this is a title"
	_obj.Overview = "this is a overview go"
	_obj.CategoryNumber = 10100
	_obj.SerialNumber = 10101

	_obj.HourRate = 15000
	_obj.WorkHours = 40

	b, _ := json.Marshal(_obj)
	_profileID := _thriftService.AddProfile(string(b))

	if !bson.IsObjectIdHex(_profileID) {
		t.Errorf("AddProfile error")
	}

	_mapUpdate := make(map[string]string)
	_mapUpdate["freelancer_name"] = "freelancer_name"
	_mapUpdate["job_title"] = "job_title"
	_mapUpdate["hour_rate"] = "1601234"
	_mapUpdate["portfolio_nums"] = "4"

	v := _thriftService.UpdateProfileBase(_profileID, _mapUpdate)

	if v != "OK" {
		t.Errorf("UpdateProfileBase error")
	}
	_profileService.Close()

}
