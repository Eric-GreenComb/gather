package command

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"testing"
)

// need start micro render service localhost:36020
func TestProfileService(t *testing.T) {

	var _profile_service ProfileService
	_thrift_service, _ := _profile_service.Default()

	_defaultObjectId := "5707cb10ae6faa1d1071a189"

	var _obj bean.Profile
	_obj.ID = bson.ObjectIdHex(_defaultObjectId)
	_obj.UserID = bson.ObjectIdHex(_defaultObjectId)
	_obj.Name = "Test"
	_obj.JobTitle = "this is a title"
	_obj.Overview = "this is a overview go"
	_obj.CategoryNumber = 10100
	_obj.SerialNumber = 10101

	_obj.HourRate = 15000
	_obj.WorkHours = 40

	b, _ := json.Marshal(_obj)
	_profile_id := _thrift_service.AddProfile(string(b))

	if !bson.IsObjectIdHex(_profile_id) {
		t.Errorf("AddProfile error")
	}

	_map_update := make(map[string]string)
	_map_update["freelancer_name"] = "freelancer_name"
	_map_update["job_title"] = "job_title"
	_map_update["hour_rate"] = "1601234"
	_map_update["portfolio_nums"] = "4"

	v := _thrift_service.UpdateProfileBase(_profile_id, _map_update)

	if v != "OK" {
		t.Errorf("UpdateProfileBase error")
	}
	_profile_service.Close()

}
