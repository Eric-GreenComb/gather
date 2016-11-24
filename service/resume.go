package service

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"gopkg.in/mgo.v2/bson"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
)

// ResumeService ResumeService
type ResumeService struct {
}

var _resumeCommandService command.ResumeService
var _resumeQueryService query.ResumeService

/**
 * command section
 */

// AddResume add resume json
func (rs *ResumeService) AddResume(resume string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.AddResume(resume)
	return
}

// AddResumeBean add resume bean
func (rs *ResumeService) AddResumeBean(resume bean.Resume) (v string) {
	b, err := json.Marshal(resume)
	if err != nil {
		return err.Error()
	}
	v = rs.AddResume(string(b))
	return
}

// UpdateResume update resume by json
func (rs *ResumeService) UpdateResume(userID string, resume string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResume(userID, resume)
	return
}

// UpdateResumeBean update resume by bean
func (rs *ResumeService) UpdateResumeBean(userID string, resume bean.Resume) (v string) {
	b, err := json.Marshal(resume)
	if err != nil {
		return err.Error()
	}
	v = rs.UpdateResume(userID, string(b))
	return
}

// UpdateResumeBase update resume base info by map
func (rs *ResumeService) UpdateResumeBase(userID string, mmap map[string]string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResumeBase(userID, mmap)
	return
}

// UpdateResumeSkillExperience update resume skill experience json
func (rs *ResumeService) UpdateResumeSkillExperience(userID string, experienceLevels string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResumeSkillExperience(userID, experienceLevels)
	return
}

// UpdateResumeSkillExperienceBean update resume skill experience bean
func (rs *ResumeService) UpdateResumeSkillExperienceBean(userID string, experienceLevels []bean.SkillExperience) (v string) {
	b, err := json.Marshal(experienceLevels)
	if err != nil {
		return err.Error()
	}
	v = rs.UpdateResumeSkillExperience(userID, string(b))
	return
}

// UpdateResumeToolandArchs update resume tool&archs json
func (rs *ResumeService) UpdateResumeToolandArchs(userID string, toolArchs string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResumeToolandArchs(userID, toolArchs)
	return
}

// UpdateResumeToolandArchsBean update resume tool&arch beans
func (rs *ResumeService) UpdateResumeToolandArchsBean(userID string, toolArchs []bean.ToolandArch) (v string) {
	b, err := json.Marshal(toolArchs)
	if err != nil {
		return err.Error()
	}
	v = rs.UpdateResumeToolandArchs(userID, string(b))
	return
}

// UpdateResumePortfolioes update resume portfolioes json
func (rs *ResumeService) UpdateResumePortfolioes(userID string, portfolioes string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResumePortfolioes(userID, portfolioes)
	return
}

// UpdateResumePortfolioesBean update resume portfolio beans
func (rs *ResumeService) UpdateResumePortfolioesBean(userID string, portfolioes []bean.Portfolio) (v string) {
	b, err := json.Marshal(portfolioes)
	if err != nil {
		return err.Error()
	}
	v = rs.UpdateResumePortfolioes(userID, string(b))
	return
}

// UpdateResumeEmploymentHistories update resume employment history json
func (rs *ResumeService) UpdateResumeEmploymentHistories(userID string, employmentHistories string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResumeEmploymentHistories(userID, employmentHistories)
	return
}

// UpdateResumeEmploymentHistoriesBean update resume employment history bean
func (rs *ResumeService) UpdateResumeEmploymentHistoriesBean(userID string, employmentHistories []bean.EmploymentHistory) (v string) {
	b, err := json.Marshal(employmentHistories)
	if err != nil {
		return err.Error()
	}
	v = rs.UpdateResumeEmploymentHistories(userID, string(b))
	return
}

// UpdateResumeEducations update resume education json
func (rs *ResumeService) UpdateResumeEducations(userID string, educations string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResumeEducations(userID, educations)
	return
}

// UpdateResumeEducationsBean update resume education beans
func (rs *ResumeService) UpdateResumeEducationsBean(userID string, educations []bean.Education) (v string) {
	b, err := json.Marshal(educations)
	if err != nil {
		return err.Error()
	}
	v = rs.UpdateResumeEducations(userID, string(b))
	return
}

// UpdateResumeOtherExperiences update resume other experiences json
func (rs *ResumeService) UpdateResumeOtherExperiences(userID string, otherExperiences string) (v string) {
	_service, _err := _resumeCommandService.Default()
	if _err != nil {
		return
	}
	defer _resumeCommandService.Close()
	v = _service.UpdateResumeOtherExperiences(userID, otherExperiences)
	return
}

// UpdateResumeOtherExperiencesBean update resume other experience beans
func (rs *ResumeService) UpdateResumeOtherExperiencesBean(userID string, otherExperiences []bean.OtherExperience) (v string) {
	b, err := json.Marshal(otherExperiences)
	if err != nil {
		return err.Error()
	}
	v = rs.UpdateResumeOtherExperiences(userID, string(b))
	return
}

/**
 * query section
 */

// GetResume get resume json by userID
func (rs *ResumeService) GetResume(userID string) (v string) {
	_service, _err := _resumeQueryService.Default()
	if _err != nil {
		return
	}
	defer _resumeQueryService.Close()
	v = _service.GetResume(userID)
	return
}

// GetResumeBean get resume bean by userID
func (rs *ResumeService) GetResumeBean(userID string) (v bean.Resume) {
	_resume := self.GetResume(userID)
	bson.Unmarshal([]byte(_resume), &v)
	return
}
