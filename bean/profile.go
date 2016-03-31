package bean

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Profile struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID bson.ObjectId `bson:"user_id" json:"user_id"`

	// select key     -1 is all
	Serialnumber   int `bson:"serial_number" json:"serial_number"`     // /gather/bean/category.go Category/SubCategory Serialnumber
	HoursBilled    int `bson:"hours_billed" json:"hours_billed"`       // /gather/global/business_const.go : Hours Billed
	AvailableHours int `bson:"available_hours" json:"available_hours"` // /gather/global/business_const.go : Available Hours
	JobSuccess     int `bson:"job_success" json:"job_success"`         // /gather/global/business_const.go : Job Success
	LastActivity   int `bson:"last_activity" json:"last_activity"`     // /gather/global/business_const.go : Last Activity
	FreelancerType int `bson:"freelancer_type" json:"freelancer_type"` // /gather/global/business_const.go : Freelancer Type
	HourlyRate     int `bson:"hourly_rate" json:"hourly_rate"`         // /gather/global/business_const.go : Hours Rate
	RegionID       int `bson:"region_id" json:"region_id"`             // /gather/global/base_const.go : Region ID

	// show
	Name              string         `bson:"freelancer_name" json:"freelancer_name"`
	JobTitle          string         `bson:"job_title" json:"job_title"`
	Overview          string         `bson:"overview" json:"overview"` // searchkey
	WorkHours         int            `bson:"work_hours" json:"work_hours"`
	PortfolioNums     int            `bson:"portfolio_nums" json:"portfolio_nums"`
	SkillTags         string         `bson:"skill_tags" json:"skill_tags"`
	AgencyHourlyRates string         `bson:"agency_hourly_rates" json:"agency_hourly_rates"`
	AgencyMembers     []AgencyMember `bson:"agency_members" json:"agency_members"`

	CreatedTime time.Time `bson:"createdtime" json:"createdtime"`
}

type AgencyMember struct {
	ProfileIds bson.ObjectId `bson:"profile_id" json:"profile_id"`
	Names      string        `bson:"profile_name" json:"profile_name"`
	JobTitles  string        `bson:"profile_jobtitle" json:"profile_jobtitle"`
	Manager    bool          `bson:"manager" json:"manager"`
	Email      string        `bson:"email" json:"email"`
	Phone      string        `bson:"phone" json:"phone"`
}
