package model

type LoginRq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DeleteRq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
