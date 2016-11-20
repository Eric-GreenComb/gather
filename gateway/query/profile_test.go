package query

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/banerwai/global/bean"
	banerwaiglobal "github.com/banerwai/global/constant"
)

// need start micro render service localhost:39050
func TestProfileDefaultService(t *testing.T) {

	var _profileService ProfileService
	_thriftService, _ := _profileService.Default()

	v := _thriftService.GetProfile("5744757a48b2b40eff000001")

	var _profile bean.Profile
	json.Unmarshal([]byte(v), &_profile)

	if len(_profile.JobTitle) == 0 {
		t.Errorf("GetProfile error")
	}
	_profileService.Close()

	_profileService.Init()
	_thriftService, _ = _profileService.Open()
	defer _profileService.Close()

	optionMap := make(map[string]int64)

	optionMap["serial_number"] = 531770282584862733

	keyMap := make(map[string]string)
	keyMap["overview"] = "go"

	v1 := _thriftService.SearchProfiles(optionMap, keyMap, time.Now().Unix(), banerwaiglobal.DefaultPageSize)

	var _profiles []bean.Profile
	json.Unmarshal([]byte(v1), &_profiles)

	if len(_profiles) == 0 {
		t.Errorf("SearchProfiles error")
	}
}
