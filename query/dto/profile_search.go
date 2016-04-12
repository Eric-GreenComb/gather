package dto

import ()

type ProfileSearchDto struct {
	SerialNumber   int    `json:"serial_number"`
	HoursBilled    int    `json:"hours_billed"`
	AvailableHours int    `json:"available_hours"`
	JobSuccess     int    `json:"job_success"`
	LastActivity   int    `json:"last_activity"`
	FreelancerType int    `json:"freelancer_type"`
	HourlyRate     int    `json:"hourly_rate"`
	RegionId       int    `json:"region_id"`
	SearchKey      string `json:"search_key"`
}
