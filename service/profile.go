package service

import (
	"encoding/json"
	"errors"
	"github.com/banerwai/global/bean"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

type ProfileService struct {
}

var _command_service command.ProfileService
var _query_service query.ProfileService

/**
 * command section
 */

func (self *ProfileService) AddProfile(json_profile string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.AddProfile(json_profile)
	return
}

func (self *ProfileService) AddProfileBean(profile bean.Profile) (v string) {
	b, err := json.Marshal(profile)
	if err != nil {
		return err.Error()
	}
	v = self.AddProfile(string(b))
	return
}

func (self *ProfileService) UpdateProfile(profile_id string, json_profile string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateProfile(profile_id, json_profile)
	return
}

func (self *ProfileService) UpdateProfileBean(profile_id string, profile bean.Profile) (v string) {
	b, err := json.Marshal(profile)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateProfile(profile_id, string(b))
	return
}

func (self *ProfileService) UpdateProfileStatus(profile_id string, status bool) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateProfileStatus(profile_id, status)
	return
}

func (self *ProfileService) UpdateProfileBase(profile_id string, mmap map[string]string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateProfileBase(profile_id, mmap)
	return
}

func (self *ProfileService) UpdateProfileAgencyMembers(profile_id string, agency_members string) (v string) {
	_service, _ := _command_service.Default()
	defer _command_service.Close()
	v = _service.UpdateProfileAgencyMembers(profile_id, agency_members)
	return
}

func (self *ProfileService) UpdateProfileAgencyMembersBean(profile_id string, agency_members []bean.AgencyMember) (v string) {
	b, err := json.Marshal(agency_members)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateProfileAgencyMembers(profile_id, string(b))
	return
}

/**
 * query section
 */

func (self *ProfileService) GetProfile(profile_id string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetProfile(profile_id)
	return
}

func (self *ProfileService) GetProfileBean(profile_id string) (bean.Profile, error) {
	var profile bean.Profile
	_json := self.GetProfile(profile_id)
	if len(_json) == 0 {
		return profile, errors.New("profile :" + profile_id + " is null")
	}

	err := json.Unmarshal([]byte(_json), &profile)
	if err != nil {
		return profile, errors.New("profile unmarshal json error")
	}
	return profile, nil
}

func (self *ProfileService) GetProfilesByUserId(user_id string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetProfilesByUserId(user_id)
	return
}

func (self *ProfileService) GetProfilesByUserIdBean(user_id string) ([]bean.Profile, error) {
	var profiles []bean.Profile

	_json := self.GetProfilesByUserId(user_id)
	if len(_json) == 0 {
		return profiles, errors.New("user :" + user_id + "'s profile is null")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}

func (self *ProfileService) GetProfilesByCategory(category_id int64, timestamp int64, pagesize int64) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetProfilesByCategory(category_id, timestamp, pagesize)
	return
}

func (self *ProfileService) GetProfilesByCategoryBean(category_id int64, timestamp int64, pagesize int64) ([]bean.Profile, error) {
	var profiles []bean.Profile

	_json := self.GetProfilesByCategoryBean(category_id, timestamp, pagesize)
	if len(_json) == 0 {
		return profiles, errors.New("category_id :" + category_id + "'s profile is null")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}

func (self *ProfileService) GetProfilesBySubCategory(subcategory_id int64, timestamp int64, pagesize int64) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetProfilesBySubCategory(subcategory_id, timestamp, pagesize)
	return
}

func (self *ProfileService) GetProfilesBySubCategoryBean(subcategory_id int64, timestamp int64, pagesize int64) ([]bean.Profile, error) {
	var profiles []bean.Profile

	_json := self.GetProfilesBySubCategoryBean(subcategory_id, timestamp, pagesize)
	if len(_json) == 0 {
		return profiles, errors.New("subcategory_id :" + subcategory_id + "'s profile is null")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}

func (self *ProfileService) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	return
}

func (self *ProfileService) SearchProfilesBean(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) ([]bean.Profile, error) {
	var profiles []bean.Profile
	_json := self.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	if len(_json) == 0 {
		return profiles, errors.New("search is not result")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}
