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

// Experience Level
const (
	ExperienceLevel_ENTRY LEVEL = iota // 0
	ExperienceLevel_INTERMEDIATE
	ExperienceLevel_EXPERT = iota // 0
)
