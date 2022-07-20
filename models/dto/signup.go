package dto

type SignupRequest struct{
	Username string `json:"username" form:"username" binding:"required,alphanum,min=1"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}