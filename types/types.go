package types

type GeneralInfo struct {
	MostActiveGroup           Group          `json:"most_active_group"`
	MostActiveTime            Time           `json:"most_active_time"`
	MostActiveMembers         MemStruct      `json:"most_active_members"`
	MostActiveMembersPerTrack []ActiveMember `json:"most_active_members_per_track"`
}

type ReadableGeneralInfo struct {
	MostActiveGroup           Group                  `json:"most_active_group"`
	MostActiveTime            Time                   `json:"most_active_time"`
	MostActiveMembers         []ReadableMem          `json:"most_active_members"`
	MostActiveMembersPerTrack []ReadableActiveMember `json:"most_active_members_per_track"`
}

type Group struct {
	Name     string `json:"group"`
	Messages int    `json:"number_of_messages"`
}

type Time struct {
	Hour    float64 `json:"hour"`
	Weekday string  `json:"weekday"`
}

type MemStruct struct {
	A int `json:"+234 706 655 0353"`
	B int `json:"+234 905 702 6031"`
	C int `json:"+234 815 671 3736"`
	D int `json:"+234 812 222 9581"`
	E int `json:"+234 814 893 5061"`
	F int `json:"+234 903 316 2311"`
	G int `json:"+234 703 206 5213"`
	H int `json:"+234 814 556 4480"`
	I int `json:"+234 810 169 4302"`
	J int `json:"+234 703 809 8061"`
}

type ReadableMem struct {
	Name     string `json:"name"`
	Messages int    `json:"number_of_messages_sent"`
}
type ActiveMember struct {
	Phone        string `json:"phone_number"`
	MessageCount int    `json:"message_count"`
	GC           string `json:"group_chat"`
}

type ReadableActiveMember struct {
	Phone        string `json:"phone_number"`
	MessageCount int    `json:"message_count"`
	GC           string `json:"group_chat"`
	Name         string `json:"name"`
}

type WrappedInfo struct {
	TotalMessages     int     `json:"number_of_messages_sent"`
	MessagePerc       float64 `json:"percentile_messages"`
	MessageRank       string  `json:"messages_top_perc"`
	TotalLinks        int     `json:"number_of_links"`
	LinkPerc          float64 `json:"percentile_links"`
	ResourceRank      string  `json:"resources_top_perc"`
	TotalQuestions    int     `json:"number_of_questions"`
	QuestionPerc      float64 `json:"percentile_questions"`
	QuestionRank      string  `json:"questions_top_perc"`
	MessageImpact     float64 `json:"message_impact_score"`
	MessageImpactPerc float64 `json:"percentile_impact"`
	MessageImpactRank string  `json:"message_impact_top_perc"`
	PeakWeekday       int     `json:"peak_weekday"`
	PeakHour          int     `json:"peak_hour"`
	PeakDayHour       string  `json:"peak_day_and_hour"`
}
