package global

// Freelancer Type
const (
	FreelancerType_Independent = iota // 0
	FreelancerType_Agency             // 1
)

// Last Activity
const (
	LastActive_Within2Weeks = iota // 0
	LastActive_Within1Month        // 1
	LastActive_Within2Months
)

// Availability
const (
	Available_True = iota // 0
	Available_False
)

// Available Hours
const (
	Available_Morethan30hrs = iota // 0
	Available_Lessthan30hrs
	Available_AsNeeded
)

// Profile Visibility
const (
	Visibility_Public = iota // 0
	Visibility_OnlyLogin
	Visibility_Private
)

// Skill Experience Level
const (
	ExperienceLevel_ENTRY = iota // 0
	ExperienceLevel_INTERMEDIATE
	ExperienceLevel_EXPERT
)

// Hours Rate
const (
	HoursRate_Lessthan50Yuan = iota // 0
	HoursRate_50_100Yuan
	HoursRate_100_150Yuan
	HoursRate_150_200Yua
	HoursRate_200_300Yua
	HoursRate_Morethan300Yua
)

// Hours Billed
const (
	HoursBilled_Atleast1Hour = iota // 0
	HoursBilled_Atleast100Hour
	HoursBilled_Morethan1000Hour
)

// Job Success
const (
	JobSuccess_80 = iota // 0
	JobSuccess_90
)
