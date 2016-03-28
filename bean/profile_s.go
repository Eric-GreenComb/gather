package bean

import ()

type Profile_Search struct {
	Id                int64  `db:"id"`
	ProfileId         string `db:"profile_id"`
	Name              string `db:"freelancer_name"`
	JobTitle          string `db:"job_title"`
	Location          string `db:"location"`
	LastActivity      int    `db:"last_activity"`
	PortfolioNums     int    `db:"portfolio_nums"`
	SkillTags         string `db:"skill_tags"`
	FreelancerType    int    `db:"freelancer_type"`
	Overview          string `db:"overview"`
	HourlyRate        int    `db:"hourly_rate"`
	AgencyHourlyRates string `bson:"agency_hourly_rates" json:"agency_hourly_rates"`
	AgencyMembers     string `bson:"agency_members" json:"agency_members"`
	WorkHours         int    `db:"work_hours"`
	AvailableHours    int    `db:"available_hours"`
	SubCategory       string `db:"subcategory_id"`
}
