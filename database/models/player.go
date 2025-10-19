package models

import (
	"gin-ayo/pkg/utils"
	"github.com/google/uuid"
)

type Player struct {
	BaseModel
	Name     string                   `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Height   int64                    `json:"height" gorm:"not null"`
	Weight   int64                    `json:"weight" gorm:"not null"`
	Number   int64                    `json:"number" gorm:"not null"`
	Position utils.PlayerPositionType `json:"position" gorm:"column:position;type:player_position_types;not null;default: 'gelandang'"`
	TeamID   uuid.UUID                `json:"team_id" gorm:"column:team_id;type:uuid;default:null"`
	Team     *Team                    `json:"team,omitempty"`
}

func (*Player) TableName() string {
	return "players"
}
