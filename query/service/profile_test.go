package service

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/banerwai/gather/query/dto"
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

	_thrift_service, _ = _profile_service.Open()
	defer _profile_service.Close()

	var profile_search_dto dto.ProfileSearchDto
	profile_search_dto.SerialNumber = 1
	profile_search_dto.AvailableHours = -1
	profile_search_dto.HoursBilled = -1
	profile_search_dto.FreelancerType = 0
	profile_search_dto.HourlyRate = 1
	profile_search_dto.HoursBilled = 1
	profile_search_dto.JobSuccess = 1
	profile_search_dto.RegionId = 1
	profile_search_dto.SearchKey = "ttt"

	_search_key, err := json.Marshal(profile_search_dto)
	if err != nil {
		fmt.Println("error:", err)
	}

	v1 := _thrift_service.SearchProfiles(string(_search_key), time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)

	if v1 != "OK" {
		t.Errorf("SearchProfiles error")
	}
}
