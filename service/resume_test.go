package service

import (
	"testing"

	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
)

func TestResumeBase(t *testing.T) {

	var _service ResumeService

	var _obj bean.Resume
	_obj.Id = bson.ObjectIdHex("5707cb10ae6faa1d1071a189")
	_obj.AuthEmail = "ministor@126.com"
	_obj.UserID = bson.ObjectIdHex("5707cb10ae6faa1d1071a189")

	_obj.Phone = "12345678901"

	var lsToolandArchs []bean.ToolandArch

	var _tool1 bean.ToolandArch
	_tool1.ToolLevel = 5
	_tool1.ToolTitle = "Java"
	lsToolandArchs = append(lsToolandArchs, _tool1)

	var _tool2 bean.ToolandArch
	_tool2.ToolLevel = 2
	_tool2.ToolTitle = "Go"
	lsToolandArchs = append(lsToolandArchs, _tool2)

	_obj.ToolandArchs = lsToolandArchs

	v := _service.AddResumeBean(_obj)
	if v != "OK" {
		t.Errorf("AddResume error")
	}

	var _tool3 bean.ToolandArch
	_tool3.ToolLevel = 1
	_tool3.ToolTitle = ".Net"
	lsToolandArchs = append(lsToolandArchs, _tool3)

	_obj.ToolandArchs = lsToolandArchs

	var lsSkillExperiences []bean.SkillExperience

	var _skill1 bean.SkillExperience
	_skill1.SkillLevel = 5
	_skill1.SkillTitle = "Java"
	lsSkillExperiences = append(lsSkillExperiences, _skill1)

	_obj.SkillExperiences = lsSkillExperiences

	userid := "5707cb10ae6faa1d1071a189"

	v = _service.UpdateResumeBean(userid, _obj)

	if v != "OK" {
		t.Errorf("UpdateResumeBean error")
	}

	_map_update := make(map[string]string)
	_map_update["auth_email"] = "ministor@126.com"
	_map_update["phone"] = "13811111111"

	v = _service.UpdateResumeBase(userid, _map_update)

	if v != "OK" {
		t.Errorf("UpdateResumeBase error")
	}

	var lsEducationes []bean.Education

	var _edu1 bean.Education
	_edu1.Degree = "degree"
	_edu1.Description = "description"
	lsEducationes = append(lsEducationes, _edu1)

	v = _service.UpdateResumeEducationsBean(userid, lsEducationes)

	_bean := _service.GetResumeBean(userid)

	if _bean.AuthEmail != "ministor@126.com" {
		t.Errorf("GetResumeBean error")
	}
	if len(_bean.ToolandArchs) != 3 {
		t.Errorf("GetResumeBean error")
	}
}
