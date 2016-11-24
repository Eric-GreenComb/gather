package service

import (
	// "encoding/json"
	"fmt"
	"github.com/banerwai/global/bean"
	banerwaiglobal "github.com/banerwai/global/constant"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestProfileAddProfile(t *testing.T) {

	var _service ProfileService

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

	_profileID := _service.AddProfileBean(_obj)

	if !bson.IsObjectIdHex(_profileID) {
		t.Errorf("AddProfile error")
	}

	// _mapUpdate := make(map[string]string)
	// _mapUpdate["freelancer_name"] = "freelancer_name"
	// _mapUpdate["job_title"] = "job_title"
	// _mapUpdate["hour_rate"] = "1601234"
	// _mapUpdate["portfolio_nums"] = "4"

	// v := _service.UpdateProfileBase(_profileID, _mapUpdate)

	// if v != "OK" {
	// 	t.Errorf("UpdateProfileBase error")
	// }

	// _bean, _ := _service.GetProfileBean(_profileID)
	// fmt.Println(_bean)

	// if len(_bean.JobTitle) == 0 {
	// 	t.Errorf("GetProfile error")
	// }

	_json := _service.GetProfilesBySubCategory(10101, time.Now().Unix(), banerwaiglobal.DefaultPageSize)
	fmt.Println(_json)

	if len(_json) != 0 {
		t.Errorf(_json)
	}
	// var profiles []bean.Profile
	// err := json.Unmarshal([]byte(_json), &profiles)
	// if err != nil {
	// 	t.Errorf("Unmarshal error")
	// }

	// if len(profiles) == 0 {
	// 	t.Errorf("len error")
	// }

	// optionMap := make(map[string]int64)

	// optionMap["serial_number"] = 10101

	// keyMmap := make(map[string]string)
	// keyMmap["overview"] = "go"

	// _json := _service.SearchProfiles(optionMap, keyMmap, time.Now().Unix(), banerwaiglobal.DefaultPageSize)
	// if len(_json) == 0 {
	// 	t.Errorf(_json)
	// }

	// var profiles []bean.Profile
	// _err := json.Unmarshal([]byte(_json), &profiles)
	// if _err != nil {
	// 	t.Errorf(_err.Error())
	// }

	// fmt.Println(len(profiles))
	// if len(profiles) == 0 {
	// 	t.Errorf("len is 0")
	// }

	// _beans, _err := _service.SearchProfilesBean(optionMap, keyMmap, time.Now().Unix(), banerwaiglobal.DefaultPageSize)

	// if _err != nil {
	// 	t.Errorf(_err.Error())
	// }

	// if len(_beans) == 0 {
	// 	t.Errorf("len is 0")
	// }
}
