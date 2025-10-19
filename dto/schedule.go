package dto

type CreateSchedule struct {
	DateTime   string `json:"date_time,omitempty" form:"date_time" binding:"required,datetime=2006-01-02 15:04"`
	HomeTeamID string `json:"home_team_id,omitempty" form:"home_team_id" binding:"required,uuid4"`
	AwayTeamID string `json:"away_team_id,omitempty" form:"away_team_id" binding:"required,uuid4"`
}

type UpdateSchedule struct {
	ID         string `json:"id,omitempty" form:"id" binding:"required,uuid4"`
	DateTime   string `json:"date_time,omitempty" form:"date_time" binding:"required,datetime=2006-01-02 15:04"`
	HomeTeamID string `json:"home_team_id,omitempty" form:"home_team_id" binding:"required,uuid4"`
	AwayTeamID string `json:"away_team_id,omitempty" form:"away_team_id" binding:"required,uuid4"`
}

type DeleteSchedule struct {
	ID string `uri:"id" binding:"required,uuid4"`
}

type DetailSchedule struct {
	ID string `uri:"id" binding:"required,uuid4"`
}
