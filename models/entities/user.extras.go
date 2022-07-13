package entities

import "gorm.io/gorm"

type UserExtras struct{
	gorm.Model
	RealName string `gorm:"index"`
	User User `gorm:"foreignKey:Username;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	
}