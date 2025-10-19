package models

type AreaVillage struct {
	BaseModel
	NamaKel string `json:"nama_kel" gorm:"type:varchar(255);not null"`
	NoProp  int64  `json:"no_prop" gorm:"not null"`
	NoKab   int64  `json:"no_kab" gorm:"not null"`
	NoKec   int64  `json:"no_kec" gorm:"not null"`
	NoKel   int64  `json:"no_kel" gorm:"not null"`
}

func (*AreaVillage) TableName() string {
	return "area_villages"
}
