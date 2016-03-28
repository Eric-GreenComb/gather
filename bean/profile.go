package bean

import (
	"labix.org/v2/mgo/bson"
)

type Profile struct {
	Id              bson.ObjectId  `bson:"_id,omitempty" json:"_id"`
	UserID          bson.ObjectId  `bson:"user_id" json:"user_id"`
	Name            string         `bson:"freelancer_name" json:"freelancer_name"`
	JobTitle        string         `bson:"job_title" json:"job_title"`
	Location        string         `bson:"location" json:"job_title"`
	SkillTags       string         `bson:"skill_tags" json:"skill_tags"`
	FreelancerType  int            `bson:"freelancer_type" json:"freelancer_type"`
	Overview        string         `bson:"overview" json:"overview"`
	HourlyRate      int            `bson:"hourly_rate" json:"hourly_rate"`
	ExperienceLevel int            `bson:"experience_level" json:"experience_level"`
	AgencyMembers   []AgencyMember `bson:"agency_members" json:"agency_members"`
}

type AgencyMember struct {
	ProfileIds bson.ObjectId `bson:"profile_id" json:"profile_id"`
	Names      string        `bson:"profile_name" json:"profile_name"`
	JobTitles  string        `bson:"profile_jobtitle" json:"profile_jobtitle"`
}
