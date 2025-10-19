package models

import (
	"github.com/google/uuid"
)

type DetailResult struct {
	BaseModel
	GoalTime  string    `json:"goal_time" gorm:"column:goal_time;type:varchar(50);not null"`
	IsPenalty bool      `json:"is_penalty" gorm:"column:is_penalty;type:boolean;not null;default: false"`
	ResultID  uuid.UUID `json:"result_id" gorm:"column:result_id;type:uuid;default:null"`
	Result    *Result   `json:"result,omitempty"`
	PlayerID  uuid.UUID `json:"player_id" gorm:"column:player_id;type:uuid;default:null"`
	Player    *Player   `json:"player,omitempty"`
}

func (*DetailResult) TableName() string {
	return "detail_results"
}
