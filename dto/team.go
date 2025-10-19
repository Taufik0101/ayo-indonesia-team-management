package dto

type CreateTeam struct {
	Name          string `json:"name,omitempty" form:"name" binding:"required"`
	Address       string `json:"address,omitempty" form:"address" binding:"required"`
	Logo          string `json:"logo,omitempty" form:"logo" binding:"required"`
	Year          int64  `json:"year,omitempty" form:"year" binding:"required,numeric"`
	ProvinceID    string `json:"province_id,omitempty" form:"province_id" binding:"required,uuid4"`
	DistrictID    string `json:"district_id,omitempty" form:"district_id" binding:"required,uuid4"`
	SubDistrictID string `json:"sub_district_id,omitempty" form:"sub_district_id" binding:"required,uuid4"`
	VillageID     string `json:"village_id,omitempty" form:"village_id" binding:"required,uuid4"`
}

type UpdateTeam struct {
	ID            string `json:"id,omitempty" form:"id" binding:"required,uuid4"`
	Name          string `json:"name,omitempty" form:"name" binding:"required"`
	Address       string `json:"address,omitempty" form:"address" binding:"required"`
	Logo          string `json:"logo,omitempty" form:"logo" binding:"required"`
	Year          int64  `json:"year,omitempty" form:"year" binding:"required,numeric"`
	ProvinceID    string `json:"province_id,omitempty" form:"province_id" binding:"required,uuid4"`
	DistrictID    string `json:"district_id,omitempty" form:"district_id" binding:"required,uuid4"`
	SubDistrictID string `json:"sub_district_id,omitempty" form:"sub_district_id" binding:"required,uuid4"`
	VillageID     string `json:"village_id,omitempty" form:"village_id" binding:"required,uuid4"`
}

type DeleteTeam struct {
	ID string `uri:"id" binding:"required,uuid4"`
}
