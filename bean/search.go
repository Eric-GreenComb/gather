package bean

import ()

type ProfileSearchCondition struct {
	// select key     -1 is all
	Serialnumber   int `json:"serial_number"`   // /gather/bean/category.go Category/SubCategory Serialnumber
	HoursBilled    int `json:"hours_billed"`    // /gather/global/business_const.go : Hours Billed
	AvailableHours int `json:"available_hours"` // /gather/global/business_const.go : Available Hours
	JobSuccess     int `json:"job_success"`     // /gather/global/business_const.go : Job Success
	LastActivity   int `json:"last_activity"`   // /gather/global/business_const.go : Last Activity
	FreelancerType int `json:"freelancer_type"` // /gather/global/business_const.go : Freelancer Type
	HourlyRate     int `json:"hourly_rate"`     // /gather/global/business_const.go : Hours Rate
	RegionID       int `json:"region_id"`       // /gather/global/base_const.go : Region ID

	SearchKey string `json:"search_key"` // searchkey
}
