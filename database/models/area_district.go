package models

type AreaDistrict struct {
	BaseModel
	NamaKab string `json:"nama_kab" gorm:"type:varchar(255);not null"`
	NoProp  int64  `json:"no_prop" gorm:"not null"`
	NoKab   int64  `json:"no_kab" gorm:"not null"`
}

func (*AreaDistrict) TableName() string {
	return "area_districts"
}
