package models

type AreaProvince struct {
	BaseModel
	NamaProp string `json:"nama_prop" gorm:"type:varchar(255);not null"`
	NoProp   int64  `json:"no_prop" gorm:"not null"`
}

func (*AreaProvince) TableName() string {
	return "area_provinces"
}
