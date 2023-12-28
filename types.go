package main

type GeneralInfo struct{}

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
