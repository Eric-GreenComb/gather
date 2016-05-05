package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestProfileService(t *testing.T) {

	var _profile_service ProfileService
	_thrift_service, _ := _profile_service.Default()

	v := _thrift_service.AddProfile("email")

	if v != "email" {
		t.Errorf("AddProfile error")
	}

	v = _thrift_service.UpdateProfile("123456")

	if v != "123456" {
		t.Errorf("UpdateProfile error")
	}
	_profile_service.Close()

	_profile_service.Init()
	_thrift_service, _ = _profile_service.Open()
	defer _profile_service.Close()
	v2 := _thrift_service.DeleteProfile("1111")
	if v2 != "1111" {
		t.Errorf("DeleteProfile error")
	}
}
