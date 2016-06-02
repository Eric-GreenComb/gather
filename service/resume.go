package service

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

type ResumeService struct {
}

var _resume_command_service command.ResumeService
var _resume_query_service query.ResumeService

/**
 * command section
 */

func (self *ResumeService) AddResume(resume string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.AddResume(resume)
	return
}

func (self *ResumeService) AddResumeBean(resume bean.Resume) (v string) {
	b, err := json.Marshal(resume)
	if err != nil {
		return err.Error()
	}
	v = self.AddResume(string(b))
	return
}

func (self *ResumeService) UpdateResume(userid string, resume string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResume(userid, resume)
	return
}

func (self *ResumeService) UpdateResumeBean(userid string, resume bean.Resume) (v string) {
	b, err := json.Marshal(resume)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateResume(userid, string(b))
	return
}

func (self *ResumeService) UpdateResumeBase(userid string, mmap map[string]string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResumeBase(userid, mmap)
	return
}

func (self *ResumeService) UpdateResumeSkillExperience(userid string, experience_levels string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResumeSkillExperience(userid, experience_levels)
	return
}

func (self *ResumeService) UpdateResumeSkillExperienceBean(userid string, experience_levels []bean.SkillExperience) (v string) {
	b, err := json.Marshal(experience_levels)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateResumeSkillExperience(userid, string(b))
	return
}

func (self *ResumeService) UpdateResumeToolandArchs(userid string, tool_archs string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResumeToolandArchs(userid, tool_archs)
	return
}

func (self *ResumeService) UpdateResumeToolandArchsBean(userid string, tool_archs []bean.ToolandArch) (v string) {
	b, err := json.Marshal(tool_archs)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateResumeToolandArchs(userid, string(b))
	return
}

func (self *ResumeService) UpdateResumePortfolioes(userid string, portfolioes string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResumePortfolioes(userid, portfolioes)
	return
}

func (self *ResumeService) UpdateResumePortfolioesBean(userid string, portfolioes []bean.Portfolio) (v string) {
	b, err := json.Marshal(portfolioes)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateResumePortfolioes(userid, string(b))
	return
}

func (self *ResumeService) UpdateResumeEmploymentHistories(userid string, employment_histories string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResumeEmploymentHistories(userid, employment_histories)
	return
}

func (self *ResumeService) UpdateResumeEmploymentHistoriesBean(userid string, employment_histories []bean.EmploymentHistory) (v string) {
	b, err := json.Marshal(employment_histories)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateResumeEmploymentHistories(userid, string(b))
	return
}

func (self *ResumeService) UpdateResumeEducations(userid string, educations string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResumeEducations(userid, educations)
	return
}

func (self *ResumeService) UpdateResumeEducationsBean(userid string, educations []bean.Education) (v string) {
	b, err := json.Marshal(educations)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateResumeEducations(userid, string(b))
	return
}

func (self *ResumeService) UpdateResumeOtherExperiences(userid string, other_experiences string) (v string) {
	_service, _err := _resume_command_service.Default()
	if _err != nil {
		return
	}
	defer _resume_command_service.Close()
	v = _service.UpdateResumeOtherExperiences(userid, other_experiences)
	return
}

func (self *ResumeService) UpdateResumeOtherExperiencesBean(userid string, other_experiences []bean.OtherExperience) (v string) {
	b, err := json.Marshal(other_experiences)
	if err != nil {
		return err.Error()
	}
	v = self.UpdateResumeOtherExperiences(userid, string(b))
	return
}

/**
 * query section
 */
func (self *ResumeService) GetResume(userid string) (v string) {
	_service, _err := _resume_query_service.Default()
	if _err != nil {
		return
	}
	defer _resume_query_service.Close()
	v = _service.GetResume(userid)
	return
}

func (self *ResumeService) GetResumeBean(userid string) (v bean.Resume) {
	_resume := self.GetResume(userid)
	bson.Unmarshal([]byte(_resume), &v)
	return
}
