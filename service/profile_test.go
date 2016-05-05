package service

import (
	banerwaiglobal "github.com/banerwai/global"
	"testing"
	"time"
)

// need start micro render service localhost:6020
func TestProfileAddProfile(t *testing.T) {

	var _service ProfileService
	v := _service.AddProfile("email")

	if v != "email" {
		t.Errorf("AddProfile error")
	}

	v = _service.UpdateProfile("123456")

	if v != "123456" {
		t.Errorf("UpdateProfile error")
	}

	v = _service.DeleteProfile("1111")
	if v != "1111" {
		t.Errorf("DeleteProfile error")
	}

	v = _service.GetProfile("1234567")

	if v != "1234567" {
		t.Errorf("GetProfile error")
	}

	option_mmap := make(map[string]int64)

	option_mmap["available_hours"] = 0
	option_mmap["job_success"] = 1
	option_mmap["freelancer_type"] = 1

	key_mmap := make(map[string]string)
	key_mmap["overview"] = "_key"

	v = _service.SearchProfiles(option_mmap, key_mmap, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)

	if v != "OK" {
		t.Errorf("SearchProfiles error")
	}
}
