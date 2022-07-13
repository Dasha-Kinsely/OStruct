package entities

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Username string `gorm:"column:username;unique_index"`
	Email string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;notNull"`
	IsAdmin bool `gorm:"column:admin;notNull"`
}