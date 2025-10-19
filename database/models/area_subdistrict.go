package models

type AreaSubDistrict struct {
	BaseModel
	NamaKec string `json:"nama_kec" gorm:"type:varchar(255);not null"`
	NoProp  int64  `json:"no_prop" gorm:"not null"`
	NoKab   int64  `json:"no_kab" gorm:"not null"`
	NoKec   int64  `json:"no_kec" gorm:"not null"`
}

func (*AreaSubDistrict) TableName() string {
	return "area_sub_districts"
}
