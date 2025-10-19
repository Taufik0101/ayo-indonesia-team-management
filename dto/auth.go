package dto

type LoginUser struct {
	Email    string `json:"email,omitempty" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
}
