package models

import (
	"github.com/google/uuid"
)

type Result struct {
	BaseModel
	ScoreHome    int64           `json:"score_home" gorm:"not null"`
	ScoreAway    int64           `json:"score_away" gorm:"not null"`
	ScheduleID   uuid.UUID       `json:"schedule_id" gorm:"column:schedule_id;type:uuid;default:null"`
	Schedule     *Schedule       `json:"schedule,omitempty"`
	WinnerTeamID uuid.UUID       `json:"winner_team_id" gorm:"column:winner_team_id;type:uuid;default:null"`
	WinnerTeam   *Team           `json:"winner_team,omitempty"`
	DetailResult []*DetailResult `json:"detail_result,omitempty"`
}

func (*Result) TableName() string {
	return "results"
}
