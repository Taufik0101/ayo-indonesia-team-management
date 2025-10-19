package dto

type CreatePlayer struct {
	Name     string `json:"name,omitempty" form:"name" binding:"required"`
	Height   int64  `json:"height,omitempty" form:"height" binding:"required,numeric"`
	Weight   int64  `json:"weight,omitempty" form:"weight" binding:"required,numeric"`
	Number   int64  `json:"number,omitempty" form:"number" binding:"required,numeric"`
	Position string `json:"position,omitempty" form:"position" binding:"required,oneof=penyerang gelandang bertahan penjaga_gawang"`
	TeamID   string `json:"team_id,omitempty" form:"team_id" binding:"required,uuid4"`
}

type UpdatePlayer struct {
	ID       string `json:"id,omitempty" form:"id" binding:"required,uuid4"`
	Name     string `json:"name,omitempty" form:"name" binding:"required"`
	Height   int64  `json:"height,omitempty" form:"height" binding:"required,numeric"`
	Weight   int64  `json:"weight,omitempty" form:"weight" binding:"required,numeric"`
	Number   int64  `json:"number,omitempty" form:"number" binding:"required,numeric"`
	Position string `json:"position,omitempty" form:"position" binding:"required,oneof=penyerang gelandang bertahan penjaga_gawang"`
	TeamID   string `json:"team_id,omitempty" form:"team_id" binding:"required,uuid4"`
}

type DeletePlayer struct {
	ID string `uri:"id" binding:"required,uuid4"`
}
