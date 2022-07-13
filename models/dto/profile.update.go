package dto

type ProfileUpdateRequest struct {
	ID int64 `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
}