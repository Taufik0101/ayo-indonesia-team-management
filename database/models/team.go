package models

import "github.com/google/uuid"

type Team struct {
	BaseModel
	Name          string           `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Logo          string           `json:"logo" gorm:"column:logo;type:text;not null"`
	Address       string           `json:"address" gorm:"column:address;type:text;not null"`
	Year          int64            `json:"year" gorm:"not null"`
	ProvinceID    uuid.UUID        `json:"province_id" gorm:"column:province_id;type:uuid;default:null"`
	Province      *AreaProvince    `json:"province,omitempty"`
	DistrictID    uuid.UUID        `json:"district_id" gorm:"column:district_id;type:uuid;default:null"`
	District      *AreaDistrict    `json:"district,omitempty"`
	SubDistrictID uuid.UUID        `json:"sub_district_id" gorm:"column:sub_district_id;type:uuid;default:null"`
	SubDistrict   *AreaSubDistrict `json:"sub_district,omitempty"`
	VillageID     uuid.UUID        `json:"village_id" gorm:"column:village_id;type:uuid;default:null"`
	Village       *AreaVillage     `json:"village,omitempty"`
}

func (*Team) TableName() string {
	return "teams"
}
