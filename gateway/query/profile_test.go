package query

import (
	"encoding/json"
	"testing"
	"time"

	banerwaiglobal "github.com/banerwai/global"
	"github.com/banerwai/global/bean"
)

// need start micro render service localhost:39050
func TestProfileDefaultService(t *testing.T) {

	var _profile_service ProfileService
	_thrift_service, _ := _profile_service.Default()

	v := _thrift_service.GetProfile("5744757a48b2b40eff000001")

	var _profile bean.Profile
	json.Unmarshal([]byte(v), &_profile)

	if len(_profile.JobTitle) == 0 {
		t.Errorf("GetProfile error")
	}
	_profile_service.Close()

	_profile_service.Init()
	_thrift_service, _ = _profile_service.Open()
	defer _profile_service.Close()

	option_mmap := make(map[string]int64)

	option_mmap["serial_number"] = 531770282584862733

	key_mmap := make(map[string]string)
	key_mmap["overview"] = "go"

	v1 := _thrift_service.SearchProfiles(option_mmap, key_mmap, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)

	var _profiles []bean.Profile
	json.Unmarshal([]byte(v1), &_profiles)

	if len(_profiles) == 0 {
		t.Errorf("SearchProfiles error")
	}
}
