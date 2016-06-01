package service

import (
	"fmt"
	banerwaiglobal "github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"testing"
	"time"
)

func TestProfileAddProfile(t *testing.T) {

	var _service ProfileService

	_defaultObjectId := "5707cb10ae6faa1d1071a189"

	var _obj bean.Profile
	_obj.Id = bson.ObjectIdHex(_defaultObjectId)
	_obj.UserID = bson.ObjectIdHex(_defaultObjectId)
	_obj.Name = "Test"
	_obj.JobTitle = "this is a title"
	_obj.Overview = "this is a overview go"
	_obj.CategoryNumber = 10100
	_obj.SerialNumber = 10101

	_obj.HourRate = 15000
	_obj.WorkHours = 40

	_profile_id := _service.AddProfileBean(_obj)

	if !bson.IsObjectIdHex(_profile_id) {
		t.Errorf("AddProfile error")
	}

	_map_update := make(map[string]string)
	_map_update["freelancer_name"] = "freelancer_name"
	_map_update["job_title"] = "job_title"
	_map_update["hour_rate"] = "1601234"
	_map_update["portfolio_nums"] = "4"

	v := _service.UpdateProfileBase(_profile_id, _map_update)

	if v != "OK" {
		t.Errorf("UpdateProfileBase error")
	}

	_bean, _ := _service.GetProfileBean(_profile_id)
	fmt.Println(_bean)

	if len(_bean.JobTitle) == 0 {
		t.Errorf("GetProfile error")
	}

	option_mmap := make(map[string]int64)

	option_mmap["serial_number"] = 10102

	key_mmap := make(map[string]string)
	key_mmap["overview"] = "go"

	_beans, _ := _service.SearchProfilesBean(option_mmap, key_mmap, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)

	if len(_beans) == 0 {
		t.Errorf("SearchProfiles error")
	}
}
