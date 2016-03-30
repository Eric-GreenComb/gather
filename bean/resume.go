package bean

import (
	"labix.org/v2/mgo/bson"
)

type Resume struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID    bson.ObjectId `bson:"user_id" json:"user_id"`
	AuthEmail string        `bson:"auth_email" json:"auth_email"`
	Phone     string        `bson:"phone" json:"phone"`

	SkillExperiences    []SkillExperience   `bson:"experience_level" json:"experience_level"`
	Portfolioes         []Portfolio         `bson:"portfolioes" json:"portfolioes"`
	EmploymentHistories []EmploymentHistory `bson:"employment_histories" json:"employment_histories"`
	Educations          []Education         `bson:"educations" json:"educations"`
	OtherExperiences    []OtherExperience   `bson:"other_experiences" json:"other_experiences"`
}

type SkillExperience struct {
	SkillTitle string `bson:"skill_title" json:"skill_title"`
	SkillLevel int    `bson:"skill_level" json:"skill_level"` // /gather/global/const.go : Skill Experience Level
}

type Portfolio struct {
	ProjectTitle  string `bson:"project_title" json:"project_title"`
	Overview      string `bson:"overview" json:"overview"`
	CategoryID    string `bson:"category_id" json:"category_id"`
	SubCategoryID string `bson:"subcategory_id" json:"subcategory_id"`
	ProjectUrl    string `bson:"project_url" json:"project_url"`
	SkillTags     string `bson:"skill_tags" json:"skill_tags"`
}

type EmploymentHistory struct {
	AgencyName  string `bson:"agency_name" json:"agency_name"`
	Lacation    string `bson:"lacation_city" json:"lacation_city"`
	WorkTitle   string `bson:"work_title" json:"work_title"`
	WorkPeriod  string `bson:"work_period" json:"work_period"`
	Description string `bson:"description" json:"description"`
}

type Education struct {
	SchoolName  string `bson:"school_name" json:"school_name"`
	StudyPeriod string `bson:"study_period" json:"study_period"`
	Degree      string `bson:"degree" json:"degree"`
	StudyArea   string `bson:"study_area" json:"study_area"`
	Description string `bson:"description" json:"description"`
}

type OtherExperience struct {
	Subject     string `bson:"subject" json:"subject"`
	Description string `bson:"description" json:"description"`
}
