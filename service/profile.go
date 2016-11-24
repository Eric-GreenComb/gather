package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/banerwai/global/bean"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

// ProfileService ProfileService
type ProfileService struct {
}

var _profileCommandService command.ProfileService
var _profileQueryService query.ProfileService

/**
 * command section
 */

// AddProfile add a profile json
func (ps *ProfileService) AddProfile(jsonProfile string) (v string) {
	_service, _err := _profileCommandService.Default()
	if _err != nil {
		return
	}
	defer _profileCommandService.Close()
	v = _service.AddProfile(jsonProfile)
	return
}

// AddProfileBean add a profile bean
func (ps *ProfileService) AddProfileBean(profile bean.Profile) (v string) {
	b, err := json.Marshal(profile)
	if err != nil {
		return err.Error()
	}
	v = ps.AddProfile(string(b))
	return
}

// UpdateProfile update profile by json
func (ps *ProfileService) UpdateProfile(profileID string, jsonProfile string) (v string) {
	_service, _err := _profileCommandService.Default()
	if _err != nil {
		return
	}
	defer _profileCommandService.Close()
	v = _service.UpdateProfile(profileID, jsonProfile)
	return
}

// UpdateProfileBean update profile by bean
func (ps *ProfileService) UpdateProfileBean(profileID string, profile bean.Profile) (v string) {
	b, err := json.Marshal(profile)
	if err != nil {
		return err.Error()
	}
	v = ps.UpdateProfile(profileID, string(b))
	return
}

// UpdateProfileStatus update profile status
func (ps *ProfileService) UpdateProfileStatus(profileID string, status bool) (v string) {
	_service, _err := _profileCommandService.Default()
	if _err != nil {
		return
	}
	defer _profileCommandService.Close()
	v = _service.UpdateProfileStatus(profileID, status)
	return
}

// UpdateProfileBase update profile by map
func (ps *ProfileService) UpdateProfileBase(profileID string, mmap map[string]string) (v string) {
	_service, _err := _profileCommandService.Default()
	if _err != nil {
		return
	}
	defer _profileCommandService.Close()
	v = _service.UpdateProfileBase(profileID, mmap)
	return
}

// UpdateProfileAgencyMembers update profile agency members json
func (ps *ProfileService) UpdateProfileAgencyMembers(profileID string, agencyMembers string) (v string) {
	_service, _err := _profileCommandService.Default()
	if _err != nil {
		return
	}
	defer _profileCommandService.Close()
	v = _service.UpdateProfileAgencyMembers(profileID, agencyMembers)
	return
}

// UpdateProfileAgencyMembersBean update profile agency member beans
func (ps *ProfileService) UpdateProfileAgencyMembersBean(profileID string, agencyMembers []bean.AgencyMember) (v string) {
	b, err := json.Marshal(agencyMembers)
	if err != nil {
		return err.Error()
	}
	v = ps.UpdateProfileAgencyMembers(profileID, string(b))
	return
}

/**
 * query section
 */

// GetProfile get profile json by ID
func (ps *ProfileService) GetProfile(profileID string) (v string) {
	_service, _err := _profileQueryService.Default()
	if _err != nil {
		return
	}
	defer _profileQueryService.Close()
	v = _service.GetProfile(profileID)
	return
}

// GetProfileBean get profile bean by ID
func (ps *ProfileService) GetProfileBean(profileID string) (bean.Profile, error) {
	var profile bean.Profile
	_json := ps.GetProfile(profileID)
	if len(_json) == 0 {
		return profile, errors.New("profile :" + profileID + " is null")
	}

	err := json.Unmarshal([]byte(_json), &profile)
	if err != nil {
		return profile, errors.New("profile unmarshal json error")
	}
	return profile, nil
}

// GetProfilesByUserID get profiles json by userID
func (ps *ProfileService) GetProfilesByUserID(userID string) (v string) {
	_service, _err := _profileQueryService.Default()
	if _err != nil {
		return
	}
	defer _profileQueryService.Close()
	v = _service.GetProfilesByUserID(userID)
	return
}

// GetProfilesByUserIDBean get profile beans by userID
func (ps *ProfileService) GetProfilesByUserIDBean(userID string) ([]bean.Profile, error) {
	var profiles []bean.Profile

	_json := ps.GetProfilesByUserID(userID)
	if len(_json) == 0 {
		return profiles, errors.New("user :" + userID + "'s profile is null")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}

// GetProfilesByCategory get profiles json by category
func (ps *ProfileService) GetProfilesByCategory(categoryID int64, timestamp int64, pagesize int64) (v string) {
	_service, _err := _profileQueryService.Default()
	if _err != nil {
		return
	}
	defer _profileQueryService.Close()
	v = _service.GetProfilesByCategory(categoryID, timestamp, pagesize)
	return
}

// GetProfilesByCategoryBean get profile beans by category
func (ps *ProfileService) GetProfilesByCategoryBean(categoryID int64, timestamp int64, pagesize int64) ([]bean.Profile, error) {
	var profiles []bean.Profile

	_json := ps.GetProfilesByCategory(categoryID, timestamp, pagesize)
	if len(_json) == 0 {
		return profiles, errors.New("categoryID :" + string(categoryID) + "'s profile is null")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}

// GetProfilesBySubCategory get profiles json by subcategory
func (ps *ProfileService) GetProfilesBySubCategory(subcategoryID int64, timestamp int64, pagesize int64) (v string) {
	_service, _err := _profileQueryService.Default()
	if _err != nil {
		return _err.Error()
	}
	defer _profileQueryService.Close()
	v = _service.GetProfilesBySubCategory(subcategoryID, timestamp, pagesize)
	fmt.Println(v)
	return
}

// GetProfilesBySubCategoryBean get profile beans by subcategory
func (ps *ProfileService) GetProfilesBySubCategoryBean(subcategoryID int64, timestamp int64, pagesize int64) ([]bean.Profile, error) {
	var profiles []bean.Profile

	_json := ps.GetProfilesBySubCategory(subcategoryID, timestamp, pagesize)
	if len(_json) == 0 {
		return profiles, errors.New("subcategoryID :" + string(subcategoryID) + "'s profile is null")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}

// SearchProfiles search profiles json by keys map
func (ps *ProfileService) SearchProfiles(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) (v string) {
	_service, _err := _profileQueryService.Default()
	if _err != nil {
		return
	}

	defer _profileQueryService.Close()
	v = _service.SearchProfiles(optionMap, keyMap, timestamp, pagesize)
	return
}

// SearchProfilesBean search profile beans by keys map
func (ps *ProfileService) SearchProfilesBean(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) ([]bean.Profile, error) {
	var profiles []bean.Profile
	_json := ps.SearchProfiles(optionMap, keyMap, timestamp, pagesize)
	if len(_json) == 0 {
		return profiles, errors.New("search is not result")
	}

	err := json.Unmarshal([]byte(_json), &profiles)
	if err != nil {
		return profiles, errors.New("profiles unmarshal json error")
	}
	return profiles, nil
}
