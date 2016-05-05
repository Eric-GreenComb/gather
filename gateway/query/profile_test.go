package query

import (
	"testing"
	"time"

	banerwaiglobal "github.com/banerwai/global"
)

// need start micro render service localhost:9050
func TestProfileDefaultService(t *testing.T) {

	var _profile_service ProfileService
	_thrift_service, _ := _profile_service.Default()

	v := _thrift_service.GetProfile("1234567")

	if v != "1234567" {
		t.Errorf("GetProfile error")
	}
	_profile_service.Close()

	_profile_service.Init()
	_thrift_service, _ = _profile_service.Open()
	defer _profile_service.Close()

	option_mmap := make(map[string]int64)

	option_mmap["available_hours"] = 0
	option_mmap["job_success"] = 1
	option_mmap["freelancer_type"] = 1

	key_mmap := make(map[string]string)
	key_mmap["overview"] = "_key"

	v1 := _thrift_service.SearchProfiles(option_mmap, key_mmap, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)

	if v1 != "OK" {
		t.Errorf("SearchProfiles error")
	}
}
