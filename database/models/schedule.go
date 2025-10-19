package models

import (
	"github.com/google/uuid"
	"time"
)

type Schedule struct {
	BaseModel
	Date         time.Time `json:"date" gorm:"column:date"`
	HomeTeamID   uuid.UUID `json:"home_team_id" gorm:"column:home_team_id;type:uuid;default:null"`
	HomeTeam     *Team     `json:"home_team,omitempty"`
	AwayTeamID   uuid.UUID `json:"away_team_id" gorm:"column:away_team_id;type:uuid;default:null"`
	AwayTeam     *Team     `json:"away_team,omitempty"`
	WinnerTeamID uuid.UUID `json:"winner_team_id" gorm:"column:winner_team_id;type:uuid;default:null"`
	WinnerTeam   *Team     `json:"winner_team,omitempty"`
	IsFinished   bool      `json:"is_finished" gorm:"column:is_finished;type:boolean;not null;default: false"`
}

func (*Schedule) TableName() string {
	return "schedules"
}
