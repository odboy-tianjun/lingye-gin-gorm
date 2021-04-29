package dto

import "lingye-gin/src/base"

type UserDTO struct {
	base.LingYeDTO
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
