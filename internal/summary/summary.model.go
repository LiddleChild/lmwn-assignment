package summary

type Summary struct {
	Province map[string]int
	AgeGroup AgeGroup
}

type AgeGroup struct {
	Young     int `json:"0-30"`
	MiddleAge int `json:"31-60"`
	Elderly   int `json:"61+"`
	Null      int `json:"N/A"`
}
